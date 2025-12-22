package customcmd

import (
	"fmt"

	"github.com/pius706975/the-sims-backend/package/database"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var name string
var username string
var email string
var password string

var CreateSuperUserCMD = &cobra.Command{
	Use:   "create-superuser",
	Short: "Create a super user",
	RunE:  createSuperUser,
}

func init() {
	CreateSuperUserCMD.Flags().StringVar(&name, "name", "", "Admin name")
	CreateSuperUserCMD.Flags().StringVar(&username, "username", "", "Admin username")
	CreateSuperUserCMD.Flags().StringVar(&email, "email", "", "Admin email")
	CreateSuperUserCMD.Flags().StringVar(&password, "password", "", "Admin password")

	CreateSuperUserCMD.MarkFlagRequired("name")
	CreateSuperUserCMD.MarkFlagRequired("username")
	CreateSuperUserCMD.MarkFlagRequired("email")
	CreateSuperUserCMD.MarkFlagRequired("password")
}

func createSuperUser(cmd *cobra.Command, args []string) error {
	if email == "" || password == "" {
		return fmt.Errorf("email and password are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	db, err := database.NewDB()
	if err != nil {
		return err
	}

	admin := models.User{
		Name:        name,
		Username:    username,
		RoleID:      nil,
		Email:       email,
		Password:    string(hashedPassword),
		IsActivated: true,
		IsSuperUser: true,
		CreatedAt:   utils.GetCurrentTime(),
	}

	if err := db.Create(&admin).Error; err != nil {
		return err
	}

	fmt.Println("Super admin created successfully")
	return nil
}

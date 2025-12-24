package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/middlewares"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
)

type userService struct {
	repo interfaces.UserRepo
}

func NewService(repo interfaces.UserRepo) *userService {
	return &userService{repo}
}

func (service *userService) UserRegistration(userData *models.User) (gin.H, int) {
	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	userData.Username = utils.GenerateUsername(userData.Email)
	userData.Password = hashedPassword
	// userData.RoleID = "f4e1855f-80a2-4ee5-a1ec-e80a9a3d3648"

	newData, err := service.repo.UserRegistration(userData)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			return gin.H{"status": 409, "message": "Email is already used"}, 409
		}
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 201, "message": "User created successfully", "data": newData}, 201
}

func (service *userService) CreateRefreshToken(userId string) (gin.H, int) {

	user, err := service.repo.GetUserById(userId)

	if err != nil {
		if err.Error() == "record not found" {
			return gin.H{"status": 404, "message": "User not found"}, 404
		}
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	tokenPayload := middlewares.TokenPayload{
		UserId:      user.ID,
		RoleId:      user.RoleID,
		Email:       user.Email,
		Username:    user.Username,
		Name:        user.Name,
		IsActivated: user.IsActivated,
		IsSuperUser: user.IsSuperUser,
	}

	jwt := middlewares.NewToken(tokenPayload, time.Hour*168)
	token, err := jwt.CreateToken()

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	expiresAt := time.Now().Add(time.Hour * 168)

	refreshToken := &models.RefreshToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: expiresAt,
	}

	newRefreshToken, err := service.repo.CreateRefreshToken(refreshToken)

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 201, "message": "Refresh token created successfully", "refresh_token": newRefreshToken.Token}, 201
}

func (service *userService) DeleteRefreshToken(userId string, refreshToken string) (gin.H, int) {
	err := service.repo.DeleteRefreshToken(userId, refreshToken)
	if err != nil {
		if err.Error() == "refresh token not found" {
			return gin.H{"status": 404, "message": "Refresh token not found"}, 404
		}
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "message": "Refresh token deleted successfully"}, 200
}

func (service *userService) ValidateRefreshToken(userId string, refreshToken string) (gin.H, int) {
	token, err := service.repo.ValidateRefreshToken(userId, refreshToken)
	if err != nil {
		if err.Error() == "record not found" {
			return gin.H{"status": 404, "message": "Refresh token not found"}, 404
		}
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "message": "Refresh token is valid", "token": token}, 200
}

func (service *userService) GetUsers() (gin.H, int) {
	users, err := service.repo.GetUsers()

	if err != nil {
		return gin.H{"status": 404, "message": err.Error()}, 404
	}

	return gin.H{"status": 200, "message": "All users fetched successfully", "data": users}, 200
}

func (service *userService) GetUserById(id string) (gin.H, int) {
	user, err := service.repo.GetUserById(id)

	if err != nil {
		return gin.H{"status": 500, "message": "Failed to retrieve user data"}, 500
	}

	return gin.H{"status": 200, "data": user}, 200
}

func (service *userService) GetUserByEmail(email string) (gin.H, int) {
	user, err := service.repo.GetUserByEmail(email)

	if err != nil {
		if err.Error() == "record not found" {
			return gin.H{"status": 404, "message": "User not found"}, 404
		}

		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "data": user}, 200
}

package interfaces

import (
	"github.com/pius706975/the-sims-backend/package/database/models"

	"github.com/gin-gonic/gin"
)

type UserRepo interface {
	UserRegistration(userData *models.User) (*models.User, error)
	CreateRefreshToken(refreshToken *models.RefreshToken) (*models.RefreshToken, error)
	ValidateRefreshToken(userId string, refreshToken string) (*models.RefreshToken, error)
	DeleteRefreshToken(userId string, refreshToken string) error

	GetUsers() (*models.Users, error)
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserService interface {
	UserRegistration(data *models.User) (gin.H, int)
	CreateRefreshToken(userId string) (gin.H, int)
	ValidateRefreshToken(userId string, refreshToken string) (gin.H, int)
	DeleteRefreshToken(userId string, refreshToken string) (gin.H, int)

	GetUsers() (gin.H, int)
	GetUserById(id string) (gin.H, int)
	GetUserByEmail(email string) (gin.H, int)
}

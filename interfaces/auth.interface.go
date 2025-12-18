package interfaces

import (
	"github.com/pius706975/the-sims-backend/package/database/models"

	"github.com/gin-gonic/gin"
)

type AuthRepo interface {
	SignIn(email string) (*models.User, error)
	CreateRefreshToken(refreshToken *models.RefreshToken) (*models.RefreshToken, error)
	DeleteRefreshTokenByUserId(userId string) error
	GetRefreshToken(token string) (*models.RefreshToken, error)
}

type AuthService interface {
	SignIn(data *models.User) (gin.H, int)
	CreateNewAccessToken(refreshToken string) (gin.H, int)
}
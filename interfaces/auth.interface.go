package interfaces

import (
	"github.com/pius706975/the-sims-backend/package/database/models"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRepo interface {
	SignIn(email string) (*models.User, error)
	CreateRefreshToken(refreshToken *models.RefreshToken) (*models.RefreshToken, error)
	DeleteRefreshTokenByUserId(userId string) error
	GetRefreshToken(token string) (*models.RefreshToken, error)
}

type AuthService interface {
	SignIn(data *models.User) (*TokenResponse, int, error)
	CreateNewAccessToken(refreshToken string) (*TokenResponse, int, error)
}

package auth

import (
	"time"

	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/middlewares"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"

	"github.com/gin-gonic/gin"
)

type authService struct {
	repo interfaces.AuthRepo
}

func NewService(repo interfaces.AuthRepo) *authService {
	return &authService{repo}
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (service *authService) SignIn(userData *models.User) (gin.H, int) {
	user, err := service.repo.SignIn(userData.Email)

	if err != nil {
		return gin.H{"status": 401, "message": "Email or password is incorrect"}, 401
	}

	if !utils.CheckPassword(userData.Password, user.Password) {
		return gin.H{"status": 401, "message": "Email or password is incorrect"}, 401
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

	jwt := middlewares.NewToken(tokenPayload, time.Minute*15)
	accessToken, err := jwt.CreateToken()

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	refreshTokenJwt := middlewares.NewToken(tokenPayload, time.Hour*168)
	refreshToken, err := refreshTokenJwt.CreateToken()

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	_ = service.repo.DeleteRefreshTokenByUserId(user.ID)

	refreshTokenModel := &models.RefreshToken{
		UserID: user.ID,
		Token:  refreshToken,
	}

	_, err = service.repo.CreateRefreshToken(refreshTokenModel)
	if err != nil {
		return gin.H{"status": 500, "message": "Failed to save refresh token"}, 500
	}

	return gin.H{"tokens": tokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}}, 200
}

func (service *authService) CreateNewAccessToken(refreshToken string) (gin.H, int) {
	tokenData, err := service.repo.GetRefreshToken(refreshToken)
	if err != nil {
		return gin.H{"status": 401, "message": "Invalid refresh token"}, 401
	}

	payloadToken, err := middlewares.DecodeRefreshToken(tokenData.Token)
	if err != nil {
		return gin.H{"status": 401, "message": "Invalid token"}, 401
	}

	tokenPayload := middlewares.TokenPayload{
		UserId:      payloadToken.UserId,
		RoleId:      payloadToken.RoleId,
		Email:       payloadToken.Email,
		Username:    payloadToken.Username,
		Name:        payloadToken.Name,
		IsActivated: payloadToken.IsActivated,
		IsSuperUser: payloadToken.IsSuperUser,
	}

	jwt := middlewares.NewToken(tokenPayload, time.Minute*15)

	accessToken, err := jwt.CreateToken()
	if err != nil {
		return gin.H{"status": 500, "message": "Failed to create access token"}, 500
	}

	return gin.H{"access_token": accessToken}, 200
}

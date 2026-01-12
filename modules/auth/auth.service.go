package auth

import (
	"errors"
	"time"

	envConfig "github.com/pius706975/the-sims-backend/config"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/middlewares"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
)

type authService struct {
	repo interfaces.AuthRepo
}

func NewService(repo interfaces.AuthRepo) *authService {
	return &authService{repo}
}

func (service *authService) SignIn(userData *models.User) (*interfaces.TokenResponse, int, error) {
	user, err := service.repo.SignIn(userData.Email)
	if err != nil {
		return nil, 401, errors.New("email or password is incorrect")
	}

	if !utils.CheckPassword(userData.Password, user.Password) {
		return nil, 401, errors.New("email or password is incorrect")
	}

	cfg := envConfig.LoadConfig()

	payload := middlewares.TokenPayload{
		UserId:      user.ID,
		RoleId:      user.RoleID,
		Email:       user.Email,
		Username:    user.Username,
		Name:        user.Name,
		IsActivated: user.IsActivated,
		IsSuperUser: user.IsSuperUser,
	}

	accessClaim := middlewares.NewAccessToken(payload, time.Minute*5)
	accessToken, err := middlewares.CreateTokenWithSecret(accessClaim, []byte(cfg.JwtSecret))
	if err != nil {
		return nil, 500, err
	}

	refreshClaim := middlewares.NewAccessToken(payload, time.Hour*168)
	refreshToken, err := middlewares.CreateTokenWithSecret(
		refreshClaim,
		[]byte(cfg.JwtRefreshTokenSecret),
	)
	if err != nil {
		return nil, 500, err
	}

	_ = service.repo.DeleteRefreshTokenByUserId(user.ID)

	if _, err := service.repo.CreateRefreshToken(&models.RefreshToken{
		UserID: user.ID,
		Token:  refreshToken,
	}); err != nil {
		return nil, 500, err
	}

	return &interfaces.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, 200, nil
}

func (service *authService) SignOut(refreshToken string) (int, error) {
	if refreshToken == "" {
		return 400, errors.New("refresh token is required")
	}

	if err := service.repo.DeleteRefreshToken(refreshToken); err != nil {
		return 500, err
	}

	return 200, nil
}

func (service *authService) CreateNewAccessToken(
	refreshToken string,
) (*interfaces.TokenResponse, int, error) {

	tokenData, err := service.repo.GetRefreshToken(refreshToken)
	if err != nil {
		return nil, 401, errors.New("invalid refresh token")
	}

	// decode refresh token to get data
	payload, err := middlewares.DecodeRefreshToken(tokenData.Token)
	if err != nil {
		return nil, 401, errors.New("invalid refresh token")
	}

	cfg := envConfig.LoadConfig()

	accessClaim := middlewares.NewAccessToken(*payload, time.Minute*5)
	accessToken, err := middlewares.CreateTokenWithSecret(
		accessClaim,
		[]byte(cfg.JwtSecret),
	)
	if err != nil {
		return nil, 500, err
	}

	return &interfaces.TokenResponse{
		AccessToken: accessToken,
	}, 200, nil
}

package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	envConfig "github.com/pius706975/the-sims-backend/config"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
)

type authController struct {
	service interfaces.AuthService
}

func NewController(service interfaces.AuthService) *authController {
	return &authController{service}
}

func (controller *authController) SignIn(ctx *gin.Context) {
	var userData models.User

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	tokens, status, err := controller.service.SignIn(&userData)
	if err != nil {
		ctx.JSON(status, gin.H{"message": err.Error()})
		return
	}

	cfg := envConfig.LoadConfig()

	ctx.SetCookie(
		"refresh_token",
		tokens.RefreshToken,
		int((time.Hour * 168).Seconds()),
		"/",
		cfg.CookieDomain,
		cfg.Mode == "production",
		true, // httpOnly true
	)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": tokens.AccessToken,
	})
}

func (controller *authController) SignOut(ctx *gin.Context) {
	cfg := envConfig.LoadConfig()

	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Refresh token not found",
		})
		return
	}

	status, err := controller.service.SignOut(refreshToken)
	if err != nil {
		ctx.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.SetCookie(
		"refresh_token",
		"",
		-1, // delete cookie
		"/",
		cfg.CookieDomain,
		cfg.Mode == "production",
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed out",
	})
}

func (controller *authController) CreateNewAccessToken(ctx *gin.Context) {

	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Refresh token not found",
		})
		return
	}

	tokenResponse, status, err := controller.service.CreateNewAccessToken(refreshToken)
	if err != nil {
		ctx.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": tokenResponse.AccessToken,
	})
}

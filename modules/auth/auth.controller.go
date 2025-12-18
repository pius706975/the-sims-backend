package auth

import (
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service interfaces.AuthService
}

func NewController(service interfaces.AuthService) *authController {
	return &authController{service}
}

// SignIn godoc
// @Summary Login as an authenticated user
// @Description Login with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userData body models.SignInRequest true "User data"
// @Success 200
// @Failure 401
// @Failure 500
// @Router /api/auth/signin [post]
func (controller *authController) SignIn(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var userData models.User

	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	responseData, status := controller.service.SignIn(&userData)

	ctx.JSON(status, responseData)
}

// CreateNewAccessToken godoc
// @Summary Create new access token
// @Description Create new access token by refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userData body models.CreateNewAccessTokenRequest true "User data"
// @Success 200
// @Failure 401
// @Failure 500
// @Router /api/auth/create-new-access-token [post]
func (controller *authController) CreateNewAccessToken(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var requestData struct {
		RefreshToken string `json:"refresh_token"`
	}

	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}

	responseData, status := controller.service.CreateNewAccessToken(requestData.RefreshToken)

	ctx.JSON(status, responseData)
}

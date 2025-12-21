package user

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service interfaces.UserService
}

func NewController(service interfaces.UserService) *userController {
	return &userController{service}
}

// UserRegistration godoc
// @Summary Register a new user
// @Description Register a new user with name, email, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param userData body models.UserRegistrationRequest true "User data"
// @Success 201
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /api/user/registration [post]
func (controller *userController) UserRegistration(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	var userData models.User

	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	_, err = govalidator.ValidateStruct(&userData)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if !utils.ValidatePassword(userData.Password) {
		ctx.JSON(400, gin.H{"message": "Password length at least 8 characters, has at least 1 uppercase letter, 1 lowercase letter, 1 number, and 1 special character."})
		return
	}

	responseData, status := controller.service.UserRegistration(&userData)

	ctx.JSON(status, responseData)
}

func (controller *userController) CreateRefreshToken(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")

	var token models.RefreshToken

	err := ctx.ShouldBindJSON(&token)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	if token.UserID == "" {
		ctx.JSON(400, gin.H{"message": "User ID is required"})
		return
	}

	responseData, status := controller.service.CreateRefreshToken(token.UserID)

	ctx.JSON(status, responseData)
}

func (controller *userController) DeleteRefreshToken(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")

	var token models.RefreshToken

	err := ctx.ShouldBindJSON(&token)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	if token.UserID == "" {
		ctx.JSON(400, gin.H{"message": "User ID is required"})
		return
	}

	if token.Token == "" {
		ctx.JSON(400, gin.H{"message": "Refresh token is required"})
		return
	}

	responseData, status := controller.service.DeleteRefreshToken(token.UserID, token.Token)

	ctx.JSON(status, responseData)
}

func (controller *userController) ValidateRefreshToken(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")

	var token models.RefreshToken

	err := ctx.ShouldBindJSON(&token)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	if token.UserID == "" {
		ctx.JSON(400, gin.H{"message": "User ID is required"})
		return
	}

	if token.Token == "" {
		ctx.JSON(400, gin.H{"message": "Refresh token is required"})
		return
	}

	responseData, status := controller.service.ValidateRefreshToken(token.UserID, token.Token)

	ctx.JSON(status, responseData)
}

// GetUsers godoc
// @Summary Get all users
// @Description Fetch all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/user/ [get]
func (controller *userController) GetUsers(ctx *gin.Context) {
	responseData, status := controller.service.GetUsers()
	ctx.JSON(status, responseData)
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Fetch the user details based on the ID provided
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 401
// @Failure 500
// @Router /api/user/{id} [get]
func (controller *userController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	responseData, status := controller.service.GetUserById(id)

	ctx.JSON(status, responseData)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Fetch the user profile based on the decoded ID from access token
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 401
// @Failure 500
// @Router /api/user/profile [get]
func (controller *userController) GetProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("id")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	responseData, status := controller.service.GetUserById(userID.(string))

	ctx.JSON(status, responseData)
}

// GetUserByEmail godoc
// @Summary Get user by email
// @Description Get user data by email
// @Tags Users
// @Accept json
// @Produce json
// @Param email query string true "User email"
// @Success 200
// @Failure 401
// @Failure 500
// @Router /api/user/get-user-by-email [get]
func (controller *userController) GetUserByEmail(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	email := ctx.Query("email")

	if email == "" {
		ctx.JSON(400, gin.H{"message": "Email is required"})
		return
	}

	responseData, status := controller.service.GetUserByEmail(email)

	ctx.JSON(status, responseData)
}

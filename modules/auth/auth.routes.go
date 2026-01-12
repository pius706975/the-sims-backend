package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/middlewares"
)

func AuthRoutes(router *gin.Engine, controller *authController, prefix string) {
	authGroup := router.Group(prefix + "/auth")
	{
		authGroup.POST("/signin", func(ctx *gin.Context) {
			controller.SignIn(ctx)
		})

		authGroup.POST("/signout", func(ctx *gin.Context) {
			controller.SignOut(ctx)
		})

		authGroup.POST("/create-new-access-token", func(ctx *gin.Context) {
			controller.CreateNewAccessToken(ctx)
		})

		authGroup.GET("/me", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.Me(ctx)
		})
	}
}

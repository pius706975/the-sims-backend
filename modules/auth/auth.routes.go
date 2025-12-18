package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.Engine, controller *authController, prefix string) {
	authGroup := router.Group(prefix + "/auth")
	{
		authGroup.POST("/signin", func(ctx *gin.Context) {
			controller.SignIn(ctx)
		})

		authGroup.POST("/create-new-access-token", func(ctx *gin.Context) {
			controller.CreateNewAccessToken(ctx)
		})
	}
}

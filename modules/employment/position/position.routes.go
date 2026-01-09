package position

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/middlewares"
)

func EmployeePositionRoutes(router *gin.Engine, controller *employeePositionController, prefix string) {
	group := router.Group(prefix + "/position")

	{
		// =====================================================
		// Position
		// =====================================================
		group.POST("/create-position", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.CreatePosition(ctx)
		})

		group.GET("/positions", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.GetPositions(ctx)
		})

		group.DELETE("/delete/:position_id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.DeletePosition(ctx)
		})
	}
}

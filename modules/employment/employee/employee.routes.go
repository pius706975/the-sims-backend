package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/middlewares"
)

func EmployeeRoutes(router *gin.Engine, controller *employeeController, prefix string) {
	employeeGroup := router.Group(prefix + "/employee")

	{
		employeeGroup.POST("/create", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.CreateEmployee(ctx)
		})
		employeeGroup.GET("/employees", func(ctx *gin.Context) {
			controller.GetEmployees(ctx)
		})
	}
}

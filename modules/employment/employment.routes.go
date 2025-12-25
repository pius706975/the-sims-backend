package employment

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/middlewares"
)

func EmploymentRoutes(router *gin.Engine, controller *employmentController, prefix string) {
	employmentGroup := router.Group(prefix + "/employment")

	{
		// Employee type
		employmentGroup.POST("/create", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.CreateEmployeeType(ctx)
		})
		employmentGroup.DELETE("/delete/:employee_type_id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.DeleteEmployeeType(ctx)
		})
		employmentGroup.GET("/employee-types", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.GetEmployeeTypes(ctx)
		})

		// Employement status
		employmentGroup.POST("/create/employment-status", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.CreateEmploymentStatus(ctx)
		})
		employmentGroup.DELETE("/delete/employment-status/:employment_status_id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.DeleteEmploymentStatus(ctx)
		})
		employmentGroup.GET("/employment-statuses", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.GetEmploymentStatuses(ctx)
		})
	}
}

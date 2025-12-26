package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/modules/employment/employment"
	"gorm.io/gorm"
)

func EmployeeModule(router *gin.Engine, db *gorm.DB, prefix string) {
	employeeRepo := NewRepo(db)
	employmentRepo := employment.NewRepo(db)

	employeeService := NewService(employeeRepo, employmentRepo)
	employeeController := NewController(employeeService)

	EmployeeRoutes(router, employeeController, prefix)
}

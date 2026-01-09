package position

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EmployeePosition(router *gin.Engine, db *gorm.DB, prefix string) {
	employeePositionRepo := NewEmployeePositionRepo(db)
	employeePositionService := NewEmployeePositionService(employeePositionRepo)
	employeePositionController := NewEmployeePositionController(employeePositionService)

	EmployeePositionRoutes(router, employeePositionController, prefix)
}

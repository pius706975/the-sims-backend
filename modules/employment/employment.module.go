package employment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EmploymentModule(router *gin.Engine, db *gorm.DB, prefix string)  {
	employmentRepo := NewRepo(db)
	employmentService := NewService(employmentRepo)
	employmentController := NewController(employmentService)

	EmploymentRoutes(router, employmentController, prefix)
}
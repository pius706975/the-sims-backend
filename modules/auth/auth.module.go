package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutesModule(router *gin.Engine, db *gorm.DB, prefix string) {
	authRepo := NewRepo(db)
	authService := NewService(authRepo)
	authController := NewController(authService)
	
	AuthRoutes(router, authController, prefix)
}
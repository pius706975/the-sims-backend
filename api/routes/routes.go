package routes

import (
	_ "github.com/pius706975/the-sims-backend/docs"
	"github.com/pius706975/the-sims-backend/modules/auth"
	"github.com/pius706975/the-sims-backend/modules/employment/employee"
	"github.com/pius706975/the-sims-backend/modules/employment/employment"
	"github.com/pius706975/the-sims-backend/modules/role"
	"github.com/pius706975/the-sims-backend/modules/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

const (
	APIPrefix = "/api"
)

func homeHandler(ctx *gin.Context) {
	ctx.JSON(404, gin.H{
		"status":  404,
		"message": "Sorry! Page not found",
	})
}

func RouteApp(router *gin.Engine, db *gorm.DB) error {
	router.GET(APIPrefix+"/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET(APIPrefix, homeHandler)

	role.RoleRoutesModule(router, db, APIPrefix)
	user.UserRoutesModule(router, db, APIPrefix)
	auth.AuthRoutesModule(router, db, APIPrefix)
	employment.EmploymentModule(router, db, APIPrefix)
	employee.EmployeeModule(router, db, APIPrefix)

	return nil
}

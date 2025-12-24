package employment

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
)

type employmentController struct {
	service interfaces.EmploymentService
}

func NewController(service interfaces.EmploymentService) *employmentController {
	return &employmentController{service}
}

// CreateEmployeeType godoc
// @Summary Create employee type
// @Description Create a new employee type with manual ID and name of the type
// @tags EmployeeTypes
// @Accept json
// @Produce json
// @Param employeeTypeData body models.CreateEmployeeTypeRequest true "EmployeeType data"
// @Param Authorization header string true "Authorization token"
// @Success 201
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /api/employment/create [post]
func (controller *employmentController) CreateEmployeeType(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	name, exists := ctx.Get("name")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var employeeTypeData models.EmployeeType

	err := ctx.ShouldBindJSON(&employeeTypeData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
	}

	_, err = govalidator.ValidateStruct(&employeeTypeData)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	responseData, status := controller.service.CreateEmployeeType(&employeeTypeData, name.(string))

	ctx.JSON(status, responseData)
}

// GetEmployeeTypes godoc
// @Summary Get all employee types
// @Description Fetch all employee types
// @Tags EmployeeTypes
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employment/employee-types [get]
func (controller *employmentController) GetEmployeeTypes(ctx *gin.Context) {
	responseData, status := controller.service.GetEmployeeTypes()
	ctx.JSON(status, responseData)
}
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
// @tags Employee Type
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

// DeleteEmployeeType godoc
// @Summary Delete an employee type
// @Description Delete employee type by id
// @Tags Employee Type
// @Accept json
// @Produce json
// @Param employee_type_id path string true "EmployeeType ID"
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employment/delete/{employee_type_id} [delete]
func (controller *employmentController) DeleteEmployeeType(ctx *gin.Context) {
	id := ctx.Param("employee_type_id")

	responseData, status := controller.service.DeleteEmployeeType(id)

	ctx.JSON(status, responseData)
}

// GetEmployeeTypes godoc
// @Summary Get all employee types
// @Description Fetch all employee types
// @Tags Employee Type
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

// ===========================================
// CreateEmploymentStatus godoc
// @Summary Create employment status
// @Description Create a new employment status with manual ID and name of the type
// @tags Employment Status
// @Accept json
// @Produce json
// @Param employmentStatusData body models.CreateEmploymentStatusRequest true "EmploymentStatus data"
// @Param Authorization header string true "Authorization token"
// @Success 201
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /api/employment/create/employment-status [post]
func (controller *employmentController) CreateEmploymentStatus(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	name, exists := ctx.Get("name")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var employmentStatusData models.EmploymentStatus

	err := ctx.ShouldBindJSON(&employmentStatusData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
	}

	_, err = govalidator.ValidateStruct(&employmentStatusData)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	responseData, status := controller.service.CreateEmploymentStatus(&employmentStatusData, name.(string))

	ctx.JSON(status, responseData)
}

// DeleteEmploymentStatus godoc
// @Summary Delete an employment statuses
// @Description Delete employment statuses by id
// @Tags Employment Status
// @Accept json
// @Produce json
// @Param employment_status_id path string true "EmploymentStatus ID"
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employment/delete/employment-status/{employment_status_id} [delete]
func (controller *employmentController) DeleteEmploymentStatus(ctx *gin.Context) {
	id := ctx.Param("employment_status_id")

	responseData, status := controller.service.DeleteEmploymentStatus(id)

	ctx.JSON(status, responseData)
}

// GetEmploymentStatuses godoc
// @Summary Get all employment statuses
// @Description Fetch all employment statuses
// @Tags Employment Status
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employment/employment-statuses [get]
func (controller *employmentController) GetEmploymentStatuses(ctx *gin.Context) {
	responseData, status := controller.service.GetEmploymentStatuses()
	ctx.JSON(status, responseData)
}

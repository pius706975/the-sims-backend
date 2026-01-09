package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
)

type employeeController struct {
	service interfaces.EmployeeService
}

func NewController(service interfaces.EmployeeService) *employeeController {
	return &employeeController{service}
}

// CreateEmployee godoc
// @Summary Create employee
// @Description Create a new employee with join_date and end_date in format YYYY-MM-DD
// @Tags Employee
// @Accept json
// @Produce json
// @Param employeeData body models.CreateEmployeeRequest true "Employee data"
// @Param Authorization header string true "Authorization token"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/employee/create [post]
func (controller *employeeController) CreateEmployee(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	name, exists := ctx.Get("name")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var req models.CreateEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errors := utils.FormatValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Validation error",
			"errors":  errors,
		})
		return
	}

	joinDate, err := utils.ParseDate(req.JoinDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid join_date format. Use YYYY-MM-DD"})
		return
	}
	endDate, err := utils.ParseDate(req.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
		return
	}

	employeeData := models.Employee{
		EmployeeNumber:     req.EmployeeNumber,
		FullName:           req.FullName,
		Gender:             req.Gender,
		BirthPlace:         req.BirthPlace,
		BirthDate:          req.BirthDate,
		Religion:           req.Religion,
		MaritalStatus:      req.MaritalStatus,
		Address:            req.Address,
		Phone:              req.Phone,
		Email:              req.Email,
		IdentifyCardNumber: req.IdentifyCardNumber,
		EmployeeTypeID:     req.EmployeeTypeID,
		EmploymentStatusID: req.EmploymentStatusID,
		JoinDate:           joinDate,
		EndDate:            endDate,
	}

	responseData, status := controller.service.CreateEmployee(&employeeData, name.(string))
	ctx.JSON(status, responseData)
}

// GetEmployees godoc
// @Summary Get all employees
// @Description Retrieve all employees
// @Tags Employee
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employee/employees [get]
func (controller *employeeController) GetEmployees(ctx *gin.Context) {
	responseData, status := controller.service.GetEmployees()
	ctx.JSON(status, responseData)
}

// GetEmployeeById godoc
// @Summary Get an employee detail
// @Description Retrieve an employee by id
// @Tags Employee
// @Accept json
// @Produce json
// @Param employee_id path string true "Employee ID"
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employee/detail/{employee_id} [get]
func (controller *employeeController) GetEmployeeById(ctx *gin.Context) {
	id := ctx.Param("employee_id")

	responseData, status := controller.service.GetEmployeeById(id)

	ctx.JSON(status, responseData)
}

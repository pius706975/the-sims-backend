package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
)

type employeeService struct {
	employeeRepo   interfaces.EmployeeRepo
	employmentRepo interfaces.EmploymentRepo
}

func NewService(
	employeeRepo interfaces.EmployeeRepo,
	employmentRepo interfaces.EmploymentRepo,
) *employeeService {
	return &employeeService{
		employeeRepo:   employeeRepo,
		employmentRepo: employmentRepo,
	}
}

func (service *employeeService) CreateEmployee(employeeData *models.Employee, decodedUsername string) (gin.H, int) {

	employeeData.CreatedAt = utils.GetCurrentTime()
	employeeData.CreatedBy = decodedUsername

	existingEmployee, err := service.employeeRepo.GetExistingEmployee(employeeData.EmployeeNumber)

	if err != nil {
		return gin.H{
			"status":  500,
			"message": err.Error(),
		}, 500
	}

	if existingEmployee != nil {
		return gin.H{
			"status":  400,
			"message": "Employee with this employee number already exists",
		}, 400
	}

	employeeType, err := service.employmentRepo.
		GetEmployeeTypeById(employeeData.EmployeeTypeID)

	if err != nil {
		return gin.H{
			"status":  500,
			"message": "Failed to validate employee type",
		}, 500
	}

	if employeeType == nil {
		return gin.H{
			"status":  400,
			"message": "employee_type_id is not valid",
		}, 400
	}

	employmentStatus, err := service.employmentRepo.
		GetEmploymentStatusById(employeeData.EmploymentStatusID)

	if err != nil {
		return gin.H{
			"status":  500,
			"message": "Failed to validate employment status",
		}, 500
	}

	if employmentStatus == nil {
		return gin.H{
			"status":  400,
			"message": "employment_status_id is not valid",
		}, 400
	}

	newData, err := service.employeeRepo.CreateEmployee(employeeData)

	if err != nil {
		return gin.H{
			"status":  400,
			"message": err.Error(),
		}, 500
	}

	return gin.H{
		"statue":  201,
		"message": "Employee created successfully",
		"data":    newData,
	}, 201
}

func (service *employeeService) GetEmployeeById(id string) (gin.H, int) {
	employee, err := service.employeeRepo.GetEmployeeById(id)

	if err != nil {
		return gin.H{
			"status": 400,
			"message": err.Error(),
		}, 400
	}

	return gin.H{
		"status": 200,
		"message": "Employee fetched successfully",
		"data": employee,
	}, 200
}

func (service *employeeService) GetEmployees() (gin.H, int) {
	employees, err := service.employeeRepo.GetEmployees()

	if err != nil {
		return gin.H{
			"status":  400,
			"message": err.Error(),
		}, 400
	}

	return gin.H{
		"status":  200,
		"message": "All employees fetched successfully",
		"data":    employees,
	}, 200
}

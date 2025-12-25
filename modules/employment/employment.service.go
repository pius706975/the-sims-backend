package employment

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
	"gorm.io/gorm"
)

type employmentService struct {
	repo interfaces.EmploymentRepo
}

func NewService(repo interfaces.EmploymentRepo) *employmentService {
	return &employmentService{repo}
}

func (service *employmentService) CreateEmployeeType(employeeTypeData *models.EmployeeType, decodedUsername string) (gin.H, int) {

	employeeTypeData.CreatedAt = utils.GetCurrentTime()
	employeeTypeData.CreatedBy = decodedUsername

	existingEmployeeType, err := service.repo.GetExistingEmployeeType(employeeTypeData.ID, employeeTypeData.EmployeeTypeName)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}
	if existingEmployeeType != nil {
		return gin.H{
			"status":  400,
			"message": "Employee type with the same ID or Name already exists",
		}, 400
	}

	newData, err := service.repo.CreateEmployeeType(employeeTypeData)
	if err != nil {
		return gin.H{
			"status":  500,
			"message": err.Error(),
		}, 500
	}

	return gin.H{
		"status":  201,
		"message": "Employee type created successfully",
		"data":    newData,
	}, 201
}

func (service *employmentService) DeleteEmployeeType(id string) (gin.H, int) {
	err := service.repo.DeleteEmployeeType(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gin.H{
				"status":  404,
				"message": "Employee type not found",
			}, 404
		}

		return gin.H{
			"status":  500,
			"message": err.Error(),
		}, 500
	}

	return gin.H{
		"status":  200,
		"message": "Employee type deleted successfully",
	}, 200
}

func (service *employmentService) GetEmployeeTypes() (gin.H, int) {
	employeeTypes, err := service.repo.GetEmployeeTypes()

	if err != nil {
		return gin.H{
			"status":  404,
			"message": err.Error(),
		}, 404
	}

	return gin.H{
		"status":  200,
		"message": "All employee types fetched successfully",
		"data":    employeeTypes,
	}, 200
}

// Employment status
func (service *employmentService) CreateEmploymentStatus(employmentStatusData *models.EmploymentStatus, decodedUsername string) (gin.H, int) {

	employmentStatusData.CreatedAt = utils.GetCurrentTime()
	employmentStatusData.CreatedBy = decodedUsername

	existingEmploymentStatus, err := service.repo.GetExistingEmploymentStatus(employmentStatusData.ID, employmentStatusData.EmploymentStatusName)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}
	if existingEmploymentStatus != nil {
		return gin.H{
			"status":  400,
			"message": "Employment status with the same ID or Name already exists",
		}, 400
	}

	newData, err := service.repo.CreateEmploymentStatus(employmentStatusData)
	if err != nil {
		return gin.H{
			"status":  500,
			"message": err.Error(),
		}, 500
	}

	return gin.H{
		"status":  201,
		"message": "Employment status created successfully",
		"data":    newData,
	}, 201
}

func (service *employmentService) DeleteEmploymentStatus(id string) (gin.H, int) {
	err := service.repo.DeleteEmploymentStatus(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gin.H{
				"status":  404,
				"message": "Employment status not found",
			}, 404
		}

		return gin.H{
			"status":  500,
			"message": err.Error(),
		}, 500
	}

	return gin.H{
		"status":  200,
		"message": "Employment status deleted successfully",
	}, 200
}

func (service *employmentService) GetEmploymentStatuses() (gin.H, int) {
	employmentStatuses, err := service.repo.GetEmploymentStatuses()

	if err != nil {
		return gin.H{
			"status":  404,
			"message": err.Error(),
		}, 404
	}

	return gin.H{
		"status":  200,
		"message": "All employment statuses fetched successfully",
		"data":    employmentStatuses,
	}, 200
}
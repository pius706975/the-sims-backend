package employment

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
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
		return  gin.H{"status": 500, "message": err.Error()}, 500
	}

	return  gin.H{"status": 201, "message": "Employee type created successfully", "data": newData}, 201
}

func (service *employmentService) GetEmployeeTypes() (gin.H, int) {
	employeeTypes, err := service.repo.GetEmployeeTypes()

	if err != nil {
		return gin.H{"status": 404, "message": err.Error()}, 404
	}

	return gin.H{"status": 200, "message": "All employee types fetched successfully", "data": employeeTypes}, 200
}
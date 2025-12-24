package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/package/database/models"
)

type EmploymentRepo interface {
	CreateEmployeeType(employeeTypeData *models.EmployeeType) (*models.EmployeeType, error)
	GetExistingEmployeeType(id, name string) (*models.EmployeeType, error)
	// DeleteEmployeeType(id string) error
	GetEmployeeTypes() (*models.EmployeeTypes, error)
	// GetEmployeeTypeById(id string) (*models.EmployeeType, error)
}

type EmploymentService interface {
	CreateEmployeeType(employeeTypeData *models.EmployeeType, decodedUsername string) (gin.H, int)
	// DeleteEmployeeType(id string) (gin.H, int)
	GetEmployeeTypes() (gin.H, int)
	// GetEmployeeTypeById(id string) (gin.H, int)
}

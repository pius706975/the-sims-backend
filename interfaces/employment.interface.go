package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/package/database/models"
)

type EmploymentRepo interface {
	CreateEmployeeType(employeeTypeData *models.EmployeeType) (*models.EmployeeType, error)
	GetExistingEmployeeType(id, name string) (*models.EmployeeType, error)
	DeleteEmployeeType(id string) error
	GetEmployeeTypes() (*models.EmployeeTypes, error)

	CreateEmploymentStatus(employmentStatusData *models.EmploymentStatus) (*models.EmploymentStatus, error)
	GetExistingEmploymentStatus(id, name string) (*models.EmploymentStatus, error)
	DeleteEmploymentStatus(id string) error
	GetEmploymentStatuses() (*models.EmploymentStatuses, error)
}

type EmploymentService interface {
	CreateEmployeeType(employeeTypeData *models.EmployeeType, decodedCreatorName string) (gin.H, int)
	DeleteEmployeeType(id string) (gin.H, int)
	GetEmployeeTypes() (gin.H, int)

	CreateEmploymentStatus(employmentStatusData *models.EmploymentStatus, decodedCreatorName string) (gin.H, int)
	DeleteEmploymentStatus(id string) (gin.H, int)
	GetEmploymentStatuses() (gin.H, int)
}

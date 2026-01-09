package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/package/database/models"
)

// =================================
// Employment
// =================================
type EmploymentRepo interface {
	// Employee Type
	CreateEmployeeType(employeeTypeData *models.EmployeeType) (*models.EmployeeType, error)
	GetExistingEmployeeType(id, name string) (*models.EmployeeType, error)
	DeleteEmployeeType(id string) error
	GetEmployeeTypes() (*models.EmployeeTypes, error)
	GetEmployeeTypeById(id string) (*models.EmployeeType, error)

	// Employment Status
	CreateEmploymentStatus(employmentStatusData *models.EmploymentStatus) (*models.EmploymentStatus, error)
	GetExistingEmploymentStatus(id, name string) (*models.EmploymentStatus, error)
	DeleteEmploymentStatus(id string) error
	GetEmploymentStatuses() (*models.EmploymentStatuses, error)
	GetEmploymentStatusById(id string) (*models.EmploymentStatus, error)
}

type EmploymentService interface {
	// Employee Type
	CreateEmployeeType(employeeTypeData *models.EmployeeType, decodedCreatorName string) (gin.H, int)
	DeleteEmployeeType(id string) (gin.H, int)
	GetEmployeeTypes() (gin.H, int)

	// Employment Status
	CreateEmploymentStatus(employmentStatusData *models.EmploymentStatus, decodedCreatorName string) (gin.H, int)
	DeleteEmploymentStatus(id string) (gin.H, int)
	GetEmploymentStatuses() (gin.H, int)
}

// =================================
// Employee
// =================================
type EmployeeRepo interface {
	CreateEmployee(employeeData *models.Employee) (*models.Employee, error)
	GetExistingEmployee(employeeNumber string) (*models.Employee, error)
	GetEmployees() ([]models.EmployeeRawResponse, error)
	GetEmployeeById(id string) (*models.EmployeeRawResponse, error)
}

type EmployeeService interface {
	CreateEmployee(employeeData *models.Employee, decodedCreatorName string) (gin.H, int)
	GetEmployees() (gin.H, int)
	GetEmployeeById(id string) (gin.H, int)
}

// =================================
// Position
// =================================
type PositionRepo interface {
	CreatePosition(positionData *models.Position) (*models.Position, error)
	GetExistingPosition(id, name string) (*models.Position, error)
	GetPositions() (*models.Positions, error)
	GetPositionById(id string) (*models.Position, error)
	DeletePosition(id string) error
}

type PositionService interface {
	CreatePosition(positionData *models.Position, decodedCreatorName string) (gin.H, int)
	GetPositions() (gin.H, int)
	DeletePosition(id string) (gin.H, int)
}

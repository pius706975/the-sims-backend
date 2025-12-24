package employment

import (
	"errors"

	"github.com/pius706975/the-sims-backend/package/database/models"
	"gorm.io/gorm"
)

type employmentRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *employmentRepo {
	return &employmentRepo{db}
}

func (repo *employmentRepo) CreateEmployeeType(data *models.EmployeeType) (*models.EmployeeType, error) {

	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *employmentRepo) GetEmployeeTypes() (*models.EmployeeTypes, error) {
	var data models.EmployeeTypes

	err := repo.db.
		Select("employee_type_id, employee_type_name, created_at, created_by, updated_at, updated_by").
		Order("created_at DESC").
		Find(&data).Error

	if err != nil {
		return nil, errors.New("Failed to get employee types")
	}

	if len(data) <= 0 {
		return nil, errors.New("Employee types is empty")
	}

	return &data, nil
}

func (repo *employmentRepo) GetExistingEmployeeType(id, name string) (*models.EmployeeType, error) {
	var existingEmployeeType models.EmployeeType

	err := repo.db.Where("employee_type_id = ? OR employee_type_name = ?", id, name).First(&existingEmployeeType).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &existingEmployeeType, nil
}

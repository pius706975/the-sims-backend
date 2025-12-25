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

func (repo *employmentRepo) DeleteEmployeeType(id string) error {
	result := repo.db.
		Where("employee_type_id = ?", id).
		Delete(&models.EmployeeType{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *employmentRepo) GetEmployeeTypes() (*models.EmployeeTypes, error) {
	var data models.EmployeeTypes

	err := repo.db.
		Select("employee_type_id, employee_type_name, created_at, created_by").
		Order("employee_type_id").
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

// Employment status
func (repo *employmentRepo) CreateEmploymentStatus(data *models.EmploymentStatus) (*models.EmploymentStatus, error) {

	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *employmentRepo) DeleteEmploymentStatus(id string) error {
	result := repo.db.
		Where("employment_status_id = ?", id).
		Delete(&models.EmploymentStatus{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *employmentRepo) GetExistingEmploymentStatus(id, name string) (*models.EmploymentStatus, error) {
	var existingEmploymentStatus models.EmploymentStatus

	err := repo.db.Where("employment_status_id = ? OR employment_status_name = ?", id, name).First(&existingEmploymentStatus).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &existingEmploymentStatus, nil
}

func (repo *employmentRepo) GetEmploymentStatuses() (*models.EmploymentStatuses, error) {
	var data models.EmploymentStatuses

	err := repo.db.
		Select("employment_status_id, employment_status_name, created_at, created_by").
		Order("employment_status_id").
		Find(&data).Error

	if err != nil {
		return nil, errors.New("Failed to get employment statuses")
	}

	if len(data) <= 0 {
		return nil, errors.New("Employment Statuses is empty")
	}

	return &data, nil
}

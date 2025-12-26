package employee

import (
	"github.com/pius706975/the-sims-backend/package/database/models"
	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *employeeRepo {
	return &employeeRepo{db}
}

func (repo *employeeRepo) CreateEmployee(data *models.Employee) (*models.Employee, error) {
	err := repo.db.Create(data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *employeeRepo) GetExistingEmployee(employeeNumber string) (*models.Employee, error) {
	var existingEmployee models.Employee

	err := repo.db.
		Where("employee_number", employeeNumber).
		First(&existingEmployee).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &existingEmployee, nil
}

package employee

import (
	"errors"

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

func (repo *employeeRepo) GetEmployees() ([]models.EmployeeRawResponse, error) {
	var data []models.EmployeeRawResponse

	query := `
		SELECT
			e.employee_id,
			e.employee_number,
			e.full_name,
			e.gender,
			e.birth_place,
			e.birth_date,
			e.religion,
			e.marital_status,
			e.address,
			e.phone,
			e.email,
			e.identify_card_number,
			et.employee_type_id,
			et.employee_type_name,
			es.employment_status_id,
			es.employment_status_name,
			e.join_date,
			e.end_date,
			e.is_activated,
			e.created_at,
			e.created_by,
			e.updated_at,
			e.updated_by
		FROM employees e
		LEFT JOIN employee_types et
			ON et.employee_type_id = e.employee_type_id
		LEFT JOIN employment_statuses es
			ON es.employment_status_id = e.employment_status_id
		ORDER BY e.join_date
	`

	if err := repo.db.Raw(query).Scan(&data).Error; err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("employee is empty")
	}

	return data, nil
}

func (repo *employeeRepo) GetEmployeeById(employeeId string) (*models.EmployeeRawResponse, error) {
	var data models.EmployeeRawResponse

	query := `
		SELECT
			e.employee_id,
			e.employee_number,
			e.full_name,
			e.gender,
			e.birth_place,
			e.birth_date,
			e.religion,
			e.marital_status,
			e.address,
			e.phone,
			e.email,
			e.identify_card_number,
			et.employee_type_id,
			et.employee_type_name,
			es.employment_status_id,
			es.employment_status_name,
			e.join_date,
			e.end_date,
			e.is_activated,
			e.created_at,
			e.created_by,
			e.updated_at,
			e.updated_by
		FROM employees e
		LEFT JOIN employee_types et
			ON et.employee_type_id = e.employee_type_id
		LEFT JOIN employment_statuses es
			ON es.employment_status_id = e.employment_status_id
		WHERE e.employee_id = ?
		ORDER BY e.join_date
	`

	if err := repo.db.Raw(query, employeeId).Scan(&data).Error; err != nil {
		return nil, err
	}

	if data.EmployeeID == "" {
		return nil, errors.New("employee not found")
	}

	return &data, nil
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

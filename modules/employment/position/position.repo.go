package position

import (
	"errors"

	"github.com/pius706975/the-sims-backend/package/database/models"
	"gorm.io/gorm"
)

type employeePositionRepo struct {
	db *gorm.DB
}

func NewEmployeePositionRepo(db *gorm.DB) *employeePositionRepo {
	return &employeePositionRepo{db}
}

// =====================================================
// Position
// =====================================================
func (repo *employeePositionRepo) CreatePosition(data *models.Position) (*models.Position, error) {

	err := repo.db.Create(data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *employeePositionRepo) DeletePosition(id string) error {
	result := repo.db.
		Where("position_id = ?", id).
		Delete(&models.Position{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *employeePositionRepo) GetPositions() (*models.Positions, error) {
	var data models.Positions

	err := repo.db.
		Select("position_id, position_name, created_at, created_by").
		Order("position_name").
		Find(&data).Error

	if err != nil {
		return nil, errors.New("Failed to retrieve positions")
	}

	if len(data) <= 0 {
		return nil, errors.New("No positions found")
	}

	return &data, nil
}

func (repo *employeePositionRepo) GetPositionById(id string) (*models.Position, error) {
	var data models.Position

	err := repo.db.
		Where("position_id = ?", id).
		First(&data).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &data, nil
}

func (repo *employeePositionRepo) GetExistingPosition(id, name string) (*models.Position, error) {
	var existingPosition models.Position

	err := repo.db.
		Where("position_id = ? OR position_name = ?", id, name).
		First(&existingPosition).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &existingPosition, nil
}

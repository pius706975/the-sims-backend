package auth

import (
	"errors"
	"github.com/pius706975/the-sims-backend/package/database/models"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *authRepo {
	return &authRepo{db}
}

func (repo *authRepo) SignIn(email string) (*models.User, error) {
	var data models.User

	result := repo.db.First(&data, "email = ?", email)
	if result.Error != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

func (repo *authRepo) CreateRefreshToken(refreshToken *models.RefreshToken) (*models.RefreshToken, error) {

	if err := repo.db.Create(refreshToken).Error; err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func (repo *authRepo) DeleteRefreshTokenByUserId(userId string) error {
	return repo.db.Where("user_id = ?", userId).Delete(&models.RefreshToken{}).Error
}

func (repo *authRepo) GetRefreshToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	result := repo.db.Where(("token = ?"), token).First(&refreshToken)
	return &refreshToken, result.Error
}
package repository

import (
	"errors"
	"github.com/hotkimho/realworld-api/models"
	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) Create(db *gorm.DB, user models.User) (int64, error) {

	err := db.Create(&user)
	if err.Error != nil {
		return 0, err.Error
	}

	return user.ID, nil
}

func (repo *userRepository) GetByEmail(db *gorm.DB, email string) (*models.User, error) {

	var user models.User

	err := db.Where("email = ?", email).First(&user)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err.Error
	}

	return &user, nil
}

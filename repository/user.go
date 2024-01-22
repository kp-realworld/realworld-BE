package repository

import (
	"github.com/hotkimho/realworld-api/models"
	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) Create(db *gorm.DB, user models.User) (int64, error) {

	result := DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

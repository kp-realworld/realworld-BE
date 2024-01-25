package repository

import (
	"errors"
	"github.com/hotkimho/realworld-api/controller/dto/user"
	"github.com/hotkimho/realworld-api/models"
	"github.com/sirupsen/logrus"
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

	return user.UserID, nil
}

func (repo *userRepository) GetByEmail(db *gorm.DB, email string) (*models.User, error) {

	var user models.User

	err := db.Where(&models.User{Email: email}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logrus.Error(err)

		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetByID(db *gorm.DB, id int64) (*models.User, error) {

	var user models.User

	err := db.Where(&models.User{UserID: id}).First(&user)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err.Error
	}

	return &user, nil
}

func (repo *userRepository) UpdateUserProfileByID(db *gorm.DB, updateRequest user.UpdateUserProfileRequestDTO, id int64) (*models.User, error) {

	var user models.User

	updateData := map[string]interface{}{}

	if updateRequest.Username != nil {
		updateData["username"] = *updateRequest.Username
	}
	if updateRequest.Bio != nil {
		updateData["bio"] = *updateRequest.Bio
	}
	if updateRequest.ProfileImage != nil {
		updateData["profile_image"] = *updateRequest.ProfileImage
	}

	err := db.Model(&models.User{UserID: id}).Updates(updateData).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

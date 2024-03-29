package repository

import (
	"context"
	"errors"
	"time"

	"github.com/hotkimho/realworld-api/types"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/hotkimho/realworld-api/controller/dto/user"
	"github.com/hotkimho/realworld-api/models"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) Create(db *gorm.DB, user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetByEmail(db *gorm.DB, email string) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	var user models.User

	err := db.WithContext(ctx).Where(&models.User{Email: email}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logrus.Error(err)

		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetByUsername(db *gorm.DB, username string) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	var user models.User

	err := db.WithContext(ctx).Where(&models.User{Username: username}).First(&user).Error
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	var user models.User

	err := db.WithContext(ctx).Where(&models.User{UserID: id}).First(&user)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err.Error
	}

	return &user, nil
}

func (repo *userRepository) UpdateUserProfileByID(db *gorm.DB, updateRequest userdto.UpdateUserProfileRequestDTO, id int64, password *string) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

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
	if updateRequest.Email != nil {
		updateData["email"] = *updateRequest.Email
	}
	if password != nil {
		updateData["password"] = password
	}

	err := db.WithContext(ctx).Model(&models.User{UserID: id}).Updates(updateData).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) UpdateUserProfileByUsernameAndUserID(db *gorm.DB, updateRequest userdto.UpdateUserProfileRequestDTO, username string, password *string, userID int64) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

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
	if updateRequest.Email != nil {
		updateData["email"] = *updateRequest.Email
	}
	if password != nil {
		updateData["password"] = password
	}

	err := db.WithContext(ctx).Model(models.User{}).
		Where(
			&models.User{
				UserID: userID,
			}).
		Updates(updateData).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) CheckEmailOrUsername(db *gorm.DB, email, username string) bool {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	var user models.User

	err := db.WithContext(ctx).Where(models.User{Email: email}).Or(models.User{Username: username}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

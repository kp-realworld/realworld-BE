package domain

import (
	"github.com/hotkimho/realworld-api/models"
)

func (u *UserDomain) Insert() (int64, error) {

	user := models.Users{
		Username: "kimho",
	}

	result := DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

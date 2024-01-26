package auth

import (
	"encoding/json"
	"github.com/hotkimho/realworld-api/controller/dto/auth"
	"github.com/hotkimho/realworld-api/models"
	"github.com/hotkimho/realworld-api/types"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignUpDTOToUser(requestDTO authdto.SignUpRequestDTO, hashedPassword string) models.User {
	return models.User{
		Username:     requestDTO.Username,
		Email:        requestDTO.Email,
		Password:     hashedPassword,
		ProfileImage: types.DEFAULT_PROFILE_IMAGE_URL,
	}
}

func UserToSignInResponseDTO(user models.User, token string) (string, error) {

	wrapper := authdto.SignInResponseWrapperDTO{
		User: authdto.SignInResponseDTO{
			UserID:   user.UserID,
			Username: user.Username,
			Token:    token,
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

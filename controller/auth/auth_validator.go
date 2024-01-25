package auth

import (
	"github.com/hotkimho/realworld-api/controller"
	"github.com/hotkimho/realworld-api/controller/dto/auth"
)

func ValidateSignUpRequestDTO(requestDTO auth.SignUpRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

func ValidateSignInRequestDTO(requestDTO auth.SignInRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

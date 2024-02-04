package auth

import (
	"github.com/hotkimho/realworld-api/controller"
	"github.com/hotkimho/realworld-api/controller/dto/auth"
)

func ValidateSignUpRequestDTO(requestDTO authdto.SignUpRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

func ValidateSignInRequestDTO(requestDTO authdto.SignInRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

func ValidateVerifyEmailRequestDTO(requestDTO authdto.VerifyEmailRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

func ValidateVerifyUsernameRequestDTO(requestDTO authdto.VerifyUsernameRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

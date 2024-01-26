package user

import (
	"github.com/hotkimho/realworld-api/controller"
	"github.com/hotkimho/realworld-api/controller/dto/user"
)

func ValidateUpdateUserProfileRequestDTO(requestDTO userdto.UpdateUserProfileRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

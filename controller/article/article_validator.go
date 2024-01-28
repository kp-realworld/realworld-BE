package article

import (
	"github.com/hotkimho/realworld-api/controller"
	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
)

func ValidateCreateArticleRequestDTO(requestDTO articledto.CreateArticleRequestDTO) error {
	return controller.Validate.Struct(requestDTO)

}

func ValidateReadArticleRequestDTO(requestDTO articledto.ReadArticleResponseDTO) error {
	return controller.Validate.Struct(requestDTO)
}

func ValidateUpdateArticleRequestDTO(requestDTO articledto.UpdateArticleRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

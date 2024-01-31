package comment

import (
	"github.com/hotkimho/realworld-api/controller"
	commentdto "github.com/hotkimho/realworld-api/controller/dto/comment"
)

func ValidateCreateCommentRequestDTO(requestDTO commentdto.CreateCommentRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

func ValidateReadCommentRequestDTO(requestDTO commentdto.ReadCommentsResponseDTO) error {
	return controller.Validate.Struct(requestDTO)
}

func ValidateUpdateCommentRequestDTO(requestDTO commentdto.UpdateCommentRequestDTO) error {
	return controller.Validate.Struct(requestDTO)
}

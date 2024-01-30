package article

import (
	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
	"github.com/hotkimho/realworld-api/models"
)

func CreateArticleRequestDTOToArticleModel(requestDTO articledto.CreateArticleRequestDTO, userID int64) models.Article {
	return models.Article{
		Title:       requestDTO.Title,
		Description: requestDTO.Description,
		Body:        requestDTO.Body,
		UserID:      userID,
	}
}

package responder

import (
	"encoding/json"
	"github.com/hotkimho/realworld-api/models"
	"net/http"

	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
)

func CreateArticleResponse(w http.ResponseWriter, requestDTO articledto.CreateArticleRequestDTO) {
	wrapper := articledto.CreateArticleResponseWrapperDTO{
		Article: articledto.CreateArticleResponseDTO{
			Title:       requestDTO.Title,
			Description: requestDTO.Description,
			Body:        requestDTO.Body,
			TagList:     requestDTO.TagList,
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func ReadArticleByIDResponse(w http.ResponseWriter, article models.Article) {

	isLiked := false
	if len(article.LikeBy) > 0 {
		isLiked = true
	}

	tagList := make([]string, 0)
	for _, tag := range article.TagList {
		tagList = append(tagList, tag.Tag)
	}
	wrapper := articledto.ReadArticleResponseWrapperDTO{
		Article: articledto.ReadArticleResponseDTO{
			Title:         article.Title,
			Description:   article.Body,
			Body:          article.Body,
			FavoriteCount: article.FavoriteCount,
			IsFavorited:   isLiked,
			TagList:       tagList,
			Author: articledto.ArticleAuthor{
				Username:     article.User.Username,
				Bio:          article.User.Bio,
				ProfileImage: article.User.ProfileImage,
			},
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

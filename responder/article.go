package responder

import (
	"encoding/json"
	"github.com/hotkimho/realworld-api/redis"
	"net/http"

	"github.com/hotkimho/realworld-api/models"

	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
)

func CreateArticleResponse(w http.ResponseWriter, article models.Article, tagList []string) {

	wrapper := articledto.CreateArticleResponseWrapperDTO{
		Article: articledto.CreateArticleResponseDTO{
			ID:            article.ID,
			Title:         article.Title,
			Description:   article.Description,
			Body:          article.Body,
			TagList:       tagList,
			IsFavorited:   false,
			FavoriteCount: article.FavoriteCount,
			Author: articledto.ArticleAuthor{
				AuthorID:     article.UserID,
				Username:     article.User.Username,
				Bio:          article.User.Bio,
				ProfileImage: article.User.ProfileImage,
			},
			CreatedAt: article.CreatedAt,
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

func ReadArticleByIDResponse(w http.ResponseWriter, article models.Article, isFollowed bool) {

	isLiked := false
	if len(article.Likes) > 0 {
		isLiked = true
	}

	tagList := make([]string, 0)
	for _, tag := range article.Tags {
		tagList = append(tagList, tag.Tag)
	}

	wrapper := articledto.ReadArticleResponseWrapperDTO{
		Article: articledto.ReadArticleResponseDTO{
			ID:            article.ID,
			Title:         article.Title,
			Description:   article.Body,
			Body:          article.Body,
			FavoriteCount: article.FavoriteCount,
			IsFavorited:   isLiked,
			TagList:       tagList,
			CreatedAt:     article.CreatedAt,
			UpdatedAt:     article.UpdatedAt,
			Author: articledto.ArticleAuthor{
				Username:     article.User.Username,
				Bio:          article.User.Bio,
				ProfileImage: article.User.ProfileImage,
				AuthorID:     article.User.UserID,
				Following:    &isFollowed,
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

func ReadArticleByOffsetResponse(w http.ResponseWriter, articles []models.Article, articleLikeCounts []models.ArticleLikeCount) {

	articleLikeMap := make(map[int64]int64)

	for _, articleCount := range articleLikeCounts {
		articleLikeMap[articleCount.ArticleID] = articleCount.Count
	}

	articleList := make([]articledto.ReadArticleResponseDTO, 0)

	for _, article := range articles {

		isLiked := false
		if len(article.Likes) > 0 {
			isLiked = true
		}

		tagList := make([]string, 0)
		for _, tag := range article.Tags {
			tagList = append(tagList, tag.Tag)
		}

		var likeCount int64
		if count, ok := articleLikeMap[article.ID]; ok {
			likeCount = count
		}

		articleList = append(articleList, articledto.ReadArticleResponseDTO{
			ID:            article.ID,
			Title:         article.Title,
			Description:   article.Body,
			Body:          article.Body,
			FavoriteCount: likeCount,
			IsFavorited:   isLiked,
			TagList:       tagList,
			CreatedAt:     article.CreatedAt,
			UpdatedAt:     article.UpdatedAt,
			Author: articledto.ArticleAuthor{
				Username:     article.User.Username,
				Bio:          article.User.Bio,
				ProfileImage: article.User.ProfileImage,
				AuthorID:     article.User.UserID,
			},
		})
	}

	wrapper := articledto.ReadArticleByOffsetResponseWrapperDTO{
		Articles: articleList,
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

func ReadArticlesByUser(w http.ResponseWriter, articles []models.Article) {

	articleList := make([]articledto.ReadArticleResponseDTO, 0)

	for _, article := range articles {

		isLiked := false
		if len(article.Likes) > 0 {
			isLiked = true
		}

		tagList := make([]string, 0)
		for _, tag := range article.Tags {
			tagList = append(tagList, tag.Tag)
		}

		articleList = append(articleList, articledto.ReadArticleResponseDTO{
			ID:            article.ID,
			Title:         article.Title,
			Description:   article.Body,
			Body:          article.Body,
			FavoriteCount: article.FavoriteCount,
			IsFavorited:   isLiked,
			TagList:       tagList,
			CreatedAt:     article.CreatedAt,
			UpdatedAt:     article.UpdatedAt,
			Author: articledto.ArticleAuthor{
				Username:     article.User.Username,
				Bio:          article.User.Bio,
				ProfileImage: article.User.ProfileImage,
				AuthorID:     article.User.UserID,
			},
		})
	}

	wrapper := articledto.ReadArticlesByUserResponseWrapperDTO{
		Articles: articleList,
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

func UpdateArticleResponse(w http.ResponseWriter, article models.Article, tagList []string) {

	isLiked := false
	if len(article.Likes) > 0 {
		isLiked = true
	}

	wrapper := articledto.UpdateArticleResponseWrapperDTO{
		Article: articledto.UpdateArticleResponseDTO{
			ID:            article.ID,
			Title:         article.Title,
			Description:   article.Body,
			Body:          article.Body,
			FavoriteCount: article.FavoriteCount,
			IsFavorited:   isLiked,
			TagList:       tagList,
			CreatedAt:     article.CreatedAt,
			UpdatedAt:     article.UpdatedAt,
			Author: articledto.ArticleAuthor{
				Username:     article.User.Username,
				Bio:          article.User.Bio,
				ProfileImage: article.User.ProfileImage,
				AuthorID:     article.User.UserID,
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

func CreateArticleLikeResponse(w http.ResponseWriter, code int, article models.Article) {

	isLiked := false
	if len(article.Likes) > 0 {
		isLiked = true
	}

	tagList := make([]string, 0)
	for _, tag := range article.Tags {
		tagList = append(tagList, tag.Tag)
	}

	likecount, err := redis.RedisManager.GetArticleLike(article.ID)
	if err != nil {
		likecount = 0
	}

	wrapper := articledto.CreateArticleLikeResponseWrapperDTO{
		Article: articledto.CreateArticleLikeResponseDTO{
			ID:            article.ID,
			Title:         article.Title,
			Description:   article.Body,
			Body:          article.Body,
			FavoriteCount: likecount,
			IsFavorited:   isLiked,
			TagList:       tagList,
			CreatedAt:     article.CreatedAt,
			Author: articledto.ArticleAuthor{
				Username:     article.User.Username,
				Bio:          article.User.Bio,
				ProfileImage: article.User.ProfileImage,
				AuthorID:     article.User.UserID,
			},
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}

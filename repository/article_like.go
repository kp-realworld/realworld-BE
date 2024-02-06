package repository

import (
	"gorm.io/gorm"

	"github.com/hotkimho/realworld-api/models"
)

type articleLikeRepository struct{}

func NewArticleLikeRepository() *articleLikeRepository {
	return &articleLikeRepository{}
}

func (repo *articleLikeRepository) Create(db *gorm.DB, articleID int64, userID int64) error {

	articleLike := models.ArticleLike{
		ArticleID: articleID,
		UserID:    userID,
	}

	err := db.Create(&articleLike).Error
	if err != nil {
		return err
	}

	return nil
}

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

func (repo *articleLikeRepository) Delete(db *gorm.DB, articleID int64, userID int64) error {

	err := db.Where("article_id = ? AND user_id = ?", articleID, userID).Delete(&models.ArticleLike{}).Error
	if err != nil {
		return err
	}
	return nil
}

// 이미 좋아요 했는지 확인
func (repo *articleLikeRepository) IsLiked(db *gorm.DB, articleID int64, userID int64) (bool, error) {

	articleLike := models.ArticleLike{
		ArticleID: articleID,
		UserID:    userID,
	}

	err := db.Where(&articleLike).First(&articleLike).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

package repository

import (
	"context"
	"github.com/hotkimho/realworld-api/types"
	"gorm.io/gorm"

	"github.com/hotkimho/realworld-api/models"
)

type articleLikeRepository struct{}

func NewArticleLikeRepository() *articleLikeRepository {
	return &articleLikeRepository{}
}

func (repo *articleLikeRepository) Create(db *gorm.DB, articleID int64, userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	articleLike := models.ArticleLike{
		ArticleID: articleID,
		UserID:    userID,
	}

	err := db.WithContext(ctx).Create(&articleLike).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *articleLikeRepository) Delete(db *gorm.DB, articleID int64, userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := db.WithContext(ctx).Where("article_id = ? AND user_id = ?", articleID, userID).Delete(&models.ArticleLike{}).Error
	if err != nil {
		return err
	}
	return nil
}

// 이미 좋아요 했는지 확인
func (repo *articleLikeRepository) IsLiked(db *gorm.DB, articleID int64, userID int64) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	articleLike := models.ArticleLike{
		ArticleID: articleID,
		UserID:    userID,
	}

	err := db.WithContext(ctx).Where(&articleLike).First(&articleLike).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

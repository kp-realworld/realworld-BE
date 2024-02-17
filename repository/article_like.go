package repository

import (
	"context"
	"github.com/hotkimho/realworld-api/redis"
	"github.com/hotkimho/realworld-api/types"
	"gorm.io/gorm"
	"time"

	"github.com/hotkimho/realworld-api/models"
)

type articleLikeRepository struct{}

func NewArticleLikeRepository() *articleLikeRepository {
	return &articleLikeRepository{}
}

func (repo *articleLikeRepository) Create(db *gorm.DB, articleID int64, userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := db.WithContext(ctx).Where("article_id = ? AND user_id = ?", articleID, userID).Delete(&models.ArticleLike{}).Error
	if err != nil {
		return err
	}

	return nil
}

// 이미 좋아요 했는지 확인
func (repo *articleLikeRepository) IsLiked(db *gorm.DB, articleID int64, userID int64) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
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

// 좋아요 유저와 summary 같이 저장
func (repo *articleLikeRepository) CreateWithTransaction(db *gorm.DB, articleID, userID int64) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	err := repo.Create(tx, articleID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = NewArticleLikeCountRepository().Create(tx, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 캐시에 설정하는게 실패 했다면 에러가 아닌 키를 삭제
	err = redis.RedisManager.IncreaseArticleLike(articleID)
	if err != nil {
		// 캐시 삭제도 에러가 발생하면 rollback
		err = redis.RedisManager.DeleteArticleLike(articleID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// 좋아요 유저와 summary 같이 삭제
func (repo *articleLikeRepository) DeleteWithTransaction(db *gorm.DB, articleID, userID int64) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	err := repo.Delete(tx, articleID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = NewArticleLikeCountRepository().Decrease(tx, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 캐시에 설정하는게 실패 했다면 에러가 아닌 키를 삭제
	err = redis.RedisManager.DecreaseArticleLike(articleID)
	if err != nil {
		// 캐시 삭제도 에러가 발생하면 rollback
		err = redis.RedisManager.DeleteArticleLike(articleID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

package repository

import (
	"context"
	"github.com/hotkimho/realworld-api/types"
	"time"

	"gorm.io/gorm"

	"github.com/hotkimho/realworld-api/models"
)

type articleTagRepository struct{}

func NewArticleTagRepository() *articleTagRepository {
	return &articleTagRepository{}
}

func (repo *articleTagRepository) Create(db *gorm.DB, articleID int64, tagList []string) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	articleTags := make([]models.ArticleTag, 0)

	for _, tag := range tagList {
		articleTags = append(articleTags, models.ArticleTag{
			ArticleID: articleID,
			Tag:       tag,
		})
	}

	err := db.WithContext(ctx).Create(&articleTags).Error
	if err != nil {
		return 0, err
	}

	return len(tagList), nil
}

func (repo *articleTagRepository) Update(db *gorm.DB, articleID int64, tagList []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := db.WithContext(ctx).Where(models.ArticleTag{ArticleID: articleID}).Delete(&models.ArticleTag{}).Error
	if err != nil {
		return err
	}

	_, err = repo.Create(db, articleID, tagList)
	if err != nil {
		return err
	}

	return nil
}

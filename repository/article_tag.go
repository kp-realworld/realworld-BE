package repository

import (
	"gorm.io/gorm"

	"github.com/hotkimho/realworld-api/models"
)

type articleTagRepository struct{}

func NewArticleTagRepository() *articleTagRepository {
	return &articleTagRepository{}
}

func (repo *articleTagRepository) Create(db *gorm.DB, articleID int64, tagList []string) (int, error) {

	articleTags := make([]models.ArticleTag, 0)

	for _, tag := range tagList {
		articleTags = append(articleTags, models.ArticleTag{
			ArticleID: articleID,
			Tag:       tag,
		})
	}

	err := db.Create(&articleTags).Error
	if err != nil {
		return 0, err
	}

	return len(tagList), nil
}

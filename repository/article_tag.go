package repository

import (
	"fmt"

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
	fmt.Println("articleTags : ", articleTags)

	err := db.Create(&articleTags).Error
	if err != nil {
		fmt.Println("error: ", err)
		return 0, err
	}

	return len(tagList), nil
}

func (repo *articleTagRepository) Update(db *gorm.DB, articleID int64, tagList []string) error {

	// 기존의 article_tag를 soft delete

	err := db.Where(models.ArticleTag{ArticleID: articleID}).Delete(&models.ArticleTag{}).Error
	if err != nil {
		return err
	}

	_, err = repo.Create(db, articleID, tagList)
	if err != nil {
		return err
	}

	return nil
}

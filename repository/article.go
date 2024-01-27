package repository

import "gorm.io/gorm"

type articleRepository struct{}

func NewArticleRepository() *articleRepository {
	return &articleRepository{}
}

func (repo *articleRepository) Create(db *gorm.DB) {

}

package repository

import (
	"errors"
	"gorm.io/gorm"

	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
	"github.com/hotkimho/realworld-api/models"
)

type articleRepository struct{}

func NewArticleRepository() *articleRepository {
	return &articleRepository{}
}

func (repo *articleRepository) Create(db *gorm.DB, requestDTO articledto.CreateArticleRequestDTO, userID int64) (*models.Article, error) {

	article := models.Article{
		Title:       requestDTO.Title,
		Description: requestDTO.Description,
		Body:        requestDTO.Body,
		UserID:      userID,
	}

	err := db.Preload("User", "user_id = ?", userID).Create(&article).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(article).Association("User").Find(&article.User)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

// 트랜잭션을 써서 article을 생성하고 article_tag를 생성하는 함수
func (repo *articleRepository) CreateWithTransaction(db *gorm.DB, requestDTO articledto.CreateArticleRequestDTO, userID int64) (*models.Article, error) {

	tx := db.Begin()

	article, err := repo.Create(tx, requestDTO, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(requestDTO.TagList) > 0 {
		_, err = NewArticleTagRepository().Create(tx, article.ID, requestDTO.TagList)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return article, nil
}

func (repo *articleRepository) GetByID(db *gorm.DB, articleID, userID int64) (*models.Article, error) {

	var article models.Article

	err := db.Model(article).
		Preload("User").
		Preload("Likes").
		Preload("Tags").
		First(&article, &models.Article{
			ID: articleID,
		}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &article, nil
}

func (repo *articleRepository) GetByOffset(db *gorm.DB, offset, limit int) ([]models.Article, error) {

	var articles []models.Article

	// id로 내림차순
	err := db.Model(&articles).
		Preload("User").
		Preload("Likes").
		Preload("Tags").
		Order("id desc").
		Offset(offset).
		Limit(limit).
		Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (repo *articleRepository) UpdateByID(db *gorm.DB, requestDTO articledto.UpdateArticleRequestDTO, articleID, userID int64) (*models.Article, error) {

	article := models.Article{
		ID:     articleID,
		UserID: userID,
	}

	updateData := map[string]interface{}{}

	if requestDTO.Title != nil {
		updateData["title"] = *requestDTO.Title
	}

	if requestDTO.Description != nil {
		updateData["description"] = *requestDTO.Description
	}

	if requestDTO.Body != nil {
		updateData["body"] = *requestDTO.Body
	}

	// 먼저 게시글이 있는지 확인
	err := db.Model(&models.Article{}).
		Where(article).
		First(&article).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	err = db.Model(article).
		//Where(&article).
		Updates(updateData).
		First(&article).
		Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (repo *articleRepository) DeleteByID(db *gorm.DB, articleID, userID int64) error {

	article := models.Article{
		ID:     articleID,
		UserID: userID,
	}
	// 먼저 게시글이 있는지 확인
	err := db.Model(&models.Article{}).
		Where(article).
		First(&article).
		Error
	if err != nil {
		return err
	}

	err = db.Where(article).
		Delete(&models.Article{}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *articleRepository) UpdateWithTransaction(db *gorm.DB, requestDTO articledto.UpdateArticleRequestDTO, userID, articleID int64) (*models.Article, error) {

	tx := db.Begin()

	article, err := repo.UpdateByID(tx, requestDTO, articleID, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	} else if article == nil {
		tx.Rollback()
		return nil, nil
	}

	err = NewArticleTagRepository().Update(tx, articleID, requestDTO.TagList)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return article, nil
}

func (repo *articleRepository) ValidateArticleOwner(db *gorm.DB, articleID, userID int64) error {

	var article models.Article

	err := db.Model(&article).Where(&models.Article{
		ID:     articleID,
		UserID: userID,
	}).First(&article).Error
	if err != nil {
		return err
	}

	return nil
}

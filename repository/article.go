package repository

import (
	"context"
	"errors"
	"github.com/hotkimho/realworld-api/types"

	"gorm.io/gorm"

	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
	"github.com/hotkimho/realworld-api/models"
)

type articleRepository struct{}

func NewArticleRepository() *articleRepository {
	return &articleRepository{}
}

func (repo *articleRepository) Create(db *gorm.DB, requestDTO articledto.CreateArticleRequestDTO, userID int64) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	article := models.Article{
		Title:       requestDTO.Title,
		Description: requestDTO.Description,
		Body:        requestDTO.Body,
		UserID:      userID,
	}

	err := db.WithContext(ctx).Create(&article).Error
	if err != nil {
		return nil, err
	}

	err = db.WithContext(ctx).Model(article).Association("User").Find(&article.User)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

// 트랜잭션을 써서 article을 생성하고 article_tag를 생성하는 함수
func (repo *articleRepository) CreateWithTransaction(db *gorm.DB, requestDTO articledto.CreateArticleRequestDTO, userID int64) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

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

	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	article := models.Article{
		ID: articleID,
	}

	query := db.WithContext(ctx).Model(article).
		Preload("User").
		Preload("Tags")

	if userID > 0 {
		query.Preload("Likes", "user_id = ?", userID)
	}

	err := query.First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &article, nil
}

func (repo *articleRepository) GetByOffset(db *gorm.DB, offset, limit int, userID int64) ([]models.Article, error) {

	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	var articles []models.Article

	// id로 내림차순
	query := db.WithContext(ctx).Debug().Model(&articles).
		Preload("User").
		Preload("Tags").
		Order("id desc")

	if userID > 0 {
		query.Preload("Likes", "user_id = ?", userID)
	}

	err := query.Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (repo *articleRepository) GetByOffsetAndTag(db *gorm.DB, offset, limit int, tag string) ([]models.Article, error) {

	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	var articles []models.Article

	query := db.WithContext(ctx).Model(&articles).
		Preload("User").
		Preload("Likes").
		Preload("Tags").
		Order("id desc")

	// tag 값이 있으면 tag만 데이터를 가져옴
	if tag != "" {
		query = query.Joins("LEFT JOIN article_tags ON articles.id = article_tags.article_id").
			Where("article_tags.tag = ?", tag)
	}

	err := query.Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (repo *articleRepository) GetByUserAndOffset(db *gorm.DB, offset, limit int, userID int64) ([]models.Article, error) {

	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	var articles []models.Article

	// id로 내림차순
	err := db.WithContext(ctx).Model(&articles).
		Preload("User").
		Preload("Likes").
		Preload("Tags").
		Where("user_id = ?", userID).
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

	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

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
	err := db.WithContext(ctx).Model(&models.Article{}).
		Where(article).
		First(&article).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	err = db.WithContext(ctx).Model(article).
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
	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	article := models.Article{
		ID:     articleID,
		UserID: userID,
	}
	// 먼저 게시글이 있는지 확인
	err := db.WithContext(ctx).Model(&models.Article{}).
		Where(article).
		First(&article).
		Error
	if err != nil {
		return err
	}

	err = db.WithContext(ctx).Where(article).
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
	ctx, cancel := context.WithTimeout(context.Background(), types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	var article models.Article

	err := db.WithContext(ctx).Model(&article).Where(&models.Article{
		ID:     articleID,
		UserID: userID,
	}).First(&article).Error
	if err != nil {
		return err
	}

	return nil
}

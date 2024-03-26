package repository

import (
	"context"
	"fmt"
	"github.com/hotkimho/realworld-api/models"
	"github.com/hotkimho/realworld-api/redis"
	"github.com/hotkimho/realworld-api/types"
	"gorm.io/gorm"
	"time"
)

type articleLikeCountRepository struct{}

func NewArticleLikeCountRepository() *articleLikeCountRepository {
	return &articleLikeCountRepository{}
}

// create
func (repo *articleLikeCountRepository) Create(db *gorm.DB, articleID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	// 이미 요약 row가 있는지 확인
	var count int64
	err := db.WithContext(ctx).
		Model(&models.ArticleLikeCount{}).
		Where("article_id = ?", articleID).
		Limit(1).Count(&count).Error
	if err != nil {
		return err
	}

	// 이미 생성된 경우
	if count > 0 {
		err = repo.Increase(db, articleID)
		if err != nil {
			return err
		}
		return nil
	}

	articleLikeCount := models.ArticleLikeCount{
		ArticleID: articleID,
		Count:     1,
	}

	err = db.WithContext(ctx).Create(&articleLikeCount).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *articleLikeCountRepository) GetByArticle(db *gorm.DB, articleID int64) (int64, error) {

	// read cache
	likeCount, err := redis.RedisManager.GetArticleLike(articleID)
	if err == nil {
		fmt.Println("cache hit1")
		return likeCount, nil
	}

	articleLikeCountObj := models.ArticleLikeCount{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err = db.WithContext(ctx).Where(models.ArticleLikeCount{
		ArticleID: articleID}).
		First(&articleLikeCountObj).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return 0, err
		}

		// 처음 조회 시, row가 없을 수 없으므로 생성
		articleLikeCountObj.ArticleID = articleID
		articleLikeCountObj.Count = 0
		err = db.WithContext(ctx).Create(&articleLikeCountObj).Error
		if err != nil {
			return 0, err
		}
	}

	err = redis.RedisManager.SetArticleLike(articleID, articleLikeCountObj.Count)
	if err != nil {
		return 0, err
	}

	fmt.Println("cache hit2")
	return articleLikeCountObj.Count, nil
}

func (repo *articleLikeCountRepository) GetByArticles(db *gorm.DB, articleIDs []int64) ([]models.ArticleLikeCount, error) {
	if len(articleIDs) == 0 {
		return []models.ArticleLikeCount{}, nil
	}

	articleLikeCountList := make([]models.ArticleLikeCount, 0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := db.WithContext(ctx).Where("article_id IN ?", articleIDs).
		Find(&articleLikeCountList).Error
	if err != nil {
		return nil, err
	}

	return articleLikeCountList, nil
}

func (repo *articleLikeCountRepository) Increase(db *gorm.DB, articleID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := db.WithContext(ctx).Model(&models.ArticleLikeCount{}).
		Where("article_id = ?", articleID).
		Update("count", gorm.Expr("count + ?", 1)).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *articleLikeCountRepository) Decrease(db *gorm.DB, articleID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := db.WithContext(ctx).Model(&models.ArticleLikeCount{}).
		Where("article_id = ?", articleID).
		Update("count", gorm.Expr("count - ?", 1)).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *articleLikeCountRepository) GetByArticlesWithRedis(db *gorm.DB, articleIDs []int64) ([]models.ArticleLikeCount, error) {
	var result []models.ArticleLikeCount

	// Redis에서 캐시된 값 조회
	fmt.Println("read ids : ", articleIDs)
	cachedCounts, missingIDs := redis.RedisManager.GetCachedArticleLikeCounts(articleIDs)
	result = append(result, cachedCounts...)

	fmt.Println("missing ids : ", missingIDs)
	if len(missingIDs) == 0 {
		return result, nil
	}

	// DB에서 누락된 articleIDs에 대한 데이터 조회
	dbCounts, err := repo.GetByArticles(db, missingIDs)
	if err != nil {
		return nil, err
	}

	// 결과 합치기
	result = append(result, dbCounts...)
	return result, nil
}

package repository

import (
	"context"
	"github.com/hotkimho/realworld-api/models"
	"github.com/hotkimho/realworld-api/types"
	"gorm.io/gorm"
	"time"
)

type followRepository struct{}

func NewFollowRepository() *followRepository {
	return &followRepository{}
}

// 팔로우를 생성하는 함수
func (repo *followRepository) Create(db *gorm.DB, followerID int64, followeeID int64) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	follow := models.Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}

	err := db.WithContext(ctx).Create(&follow).Error
	if err != nil {
		return err
	}

	return nil
}

// 팔로우를 삭제하는 함수
func (repo *followRepository) Delete(db *gorm.DB, followerID int64, followeeID int64) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := db.WithContext(ctx).Where("follower_id = ? AND followee_id = ?", followerID, followeeID).Delete(&models.Follow{}).Error
	if err != nil {
		return err
	}
	return nil
}

// 팔로우를 했는지 확인하는 함수
func (repo *followRepository) IsFollowing(db *gorm.DB, followerID int64, followeeID int64) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	follow := models.Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}

	err := db.WithContext(ctx).Where(&follow).First(&follow).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

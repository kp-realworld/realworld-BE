package redis

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/hotkimho/realworld-api/types"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisManager struct {
	redisClient *redis.Client
}

var RedisManager redisManager

func Init() error {
	RedisManager.Connect()

	return RedisManager.Ping()
}

// redis 연결
func (r *redisManager) Connect() {
	r.redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// redis 연결 해제
func (r *redisManager) Disconnect() {
	_ = r.redisClient.Close()
}

// health check
func (r *redisManager) Ping() error {
	// 5초 timeout context

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	_, err := r.redisClient.Ping(ctx).Result()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	//r.redisClient.Set(ctx, "health", "ok", 0)
	a, _ := r.redisClient.Get(ctx, "health").Result()
	fmt.Println(a)
	return nil
}

func (r *redisManager) GetArticleLike(articleId int64) (int64, error) {
	key := fmt.Sprintf("article:%d:like", articleId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	likeCount, err := r.redisClient.Get(ctx, key).Int64()
	if err != nil {
		// 네트워크 문제가 생긴 경우 sentry 에러로 처리
		if err != redis.Nil {
			sentry.CaptureException(err)
		}
		return 0, err
	}

	fmt.Println("cache hit , likeCount : ", likeCount)
	return likeCount, nil
}

func (r *redisManager) SetArticleLike(articleId, likeCount int64) error {
	key := fmt.Sprintf("article:%d:like", articleId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := r.redisClient.Set(ctx, key, likeCount, time.Hour*24).Err()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (r *redisManager) DeleteArticleLike(articleId int64) error {
	key := fmt.Sprintf("article:%d:like", articleId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	err := r.redisClient.Del(ctx, key).Err()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (r *redisManager) IncreaseArticleLike(articleId int64) error {
	key := fmt.Sprintf("article:%d:like", articleId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	_, err := r.redisClient.Incr(ctx, key).Result()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (r *redisManager) DecreaseArticleLike(articleId int64) error {
	key := fmt.Sprintf("article:%d:like", articleId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*types.DEFAULT_TIMEOUT_SEC)
	defer cancel()

	_, err := r.redisClient.Decr(ctx, key).Result()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

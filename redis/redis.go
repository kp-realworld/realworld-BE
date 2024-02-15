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

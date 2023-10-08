package redis

import "github.com/redis/go-redis/v9"

type RedisManager struct {
}

var RDB *redis.Client

func NewRedisConnection() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	RDB.Ping().Result()

}

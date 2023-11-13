package redist

import (
	"time"

	"github.com/go-redis/redis/v8"
)

const Expiration = 1 * time.Hour

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func CloseRedisConnection(client *redis.Client) {
	client.Close()
}

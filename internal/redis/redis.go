package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

var ctx = context.Background()

func NewRedisClient() (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
	})

	_, err := rds.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	RedisClient = rds
	return rds, nil
}

func CloseRedis(client *redis.Client) {
	err := client.Close()
	if err != nil {
		log.Fatal(err)
	}
}

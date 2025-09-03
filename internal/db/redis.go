package db

import (
	"context"
	"os"
	"github.com/redis/go-redis/v9"
	"log"
)

var RedisClient *redis.Client
var bgCtx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "redis:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})
	// Test Redis Connection
	_,err := RedisClient.Ping(bgCtx).Result()
	if err != nil {
		log.Printf("failed to connect Redis : %v", err)
	}
	log.Println("Connected to Redis")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

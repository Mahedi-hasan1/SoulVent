package repository

import (
	"context"
	"log"
	"soulvent/internal/db"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetSeenPostIDs(userID string) ([]string, error) {
	ctx := context.Background()
    key := "user:" + userID + ":viewed_posts"
    
    return db.RedisClient.ZRevRange(ctx, key, 0, -1).Result()
}

func MarkPostSeen(userID, postID string) error {
    ctx := context.Background()
    key := "user:" + userID + ":viewed_posts"
    now := float64(time.Now().Unix())
    log.Println("marking Post as Viewed :", postID)
    return db.RedisClient.ZAdd(ctx, key, redis.Z{
        Score:  now,
        Member: postID,
    }).Err()
}

func  ClearOldSeenPost(userID string, beforeTime int64) error {
    ctx := context.Background()
    key := "user:" + userID + ":viewed_posts"
    
    return db.RedisClient.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(beforeTime, 10)).Err()
}


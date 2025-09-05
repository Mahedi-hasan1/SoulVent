package repository

import (
	"fmt"
	"log"
	"soulvent/internal/db"
	"strconv"
	"time"
)

func GetSeenPostIDs(userID string) ([]string, error) {
    key := "user:" + userID + ":viewed_posts"
    
    res, err := db.UpstashRequest("ZREVRANGE", key, "0", "-1")
    if err != nil {
        return nil, err
    }
    resultMap, ok := res.(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("expected map response from UpstashRequest, got %T", res)
    }

    // Step 2: Extract the 'result' field
    result := resultMap["result"]
    if result == nil {
        return nil, fmt.Errorf("missing 'result' in response")
    }

    // Step 3: Assert that result is []interface{}
    values, ok := result.([]interface{})
    if !ok {
        return nil, fmt.Errorf("expected 'result' to be an array, got %T", result)
    }

    // Step 4: Convert []interface{} → []string
    var postIDs []string
    for _, v := range values {
        if postID, valid := v.(string); valid {
            postIDs = append(postIDs, postID)
        }
    }

    return postIDs, nil
}

func MarkPostSeen(userID, postID string) error {
    key := "user:" + userID + ":viewed_posts"
    now := float64(time.Now().Unix())
    log.Println("marking Post as Viewed :", postID)
    res, err := db.UpstashRequest("ZADD", key, now, postID)
    if err != nil {
        return fmt.Errorf("failed to mark post as seen: %w", err)
    }

    // Extract the actual result (from your wrapped map)
    resultMap, ok := res.(map[string]interface{})
    if !ok {
        return fmt.Errorf("expected map response from UpstashRequest, got %T", res)
    }

    result := resultMap["result"]
    if result == nil {
        return fmt.Errorf("missing 'result' in response")
    }

    // ✅ For ZADD, Upstash returns a number: number of elements added (as float64)
    switch v := result.(type) {
    case float64, float32, int, int64, int32:
        log.Printf("ZADD added %v elements", v)
        return nil
    default:
        return fmt.Errorf("unexpected 'result' type from ZADD: %T, value: %v", v, v)
    }
}

func  ClearOldSeenPost(userID string, beforeTime int64) error {
    key := "user:" + userID + ":viewed_posts"
    
    scoreLimit := strconv.FormatInt(beforeTime, 10)

    res, err := db.UpstashRequest("ZREMRANGEBYSCORE", key, "0", scoreLimit)
    if err != nil {
        return fmt.Errorf("failed to clear old seen posts: %w", err)
    }

    // Step 1: Extract the actual result from the wrapper map (if applicable)
    resultMap, ok := res.(map[string]interface{})
    if !ok {
        return fmt.Errorf("expected map response from UpstashRequest, got %T", res)
    }

    result := resultMap["result"]
    if result == nil {
        return fmt.Errorf("missing 'result' in response")
    }
    return nil
    // Step 2: ZREMRANGEBYSCORE returns a number (number of deleted elements)
    // switch v := result.(type) {
    // case float64, float32, int, int64, int32:
    //     deletedCount := int(v)
    //     log.Printf("Cleared %d old seen posts for user %s", deletedCount, userID)
    //     return nil // Success
    // default:
    //     return fmt.Errorf("unexpected 'result' type from ZREMRANGEBYSCORE: %T, value: %v", v, v)
    // }
}


// db/redis.go
package db

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var httpClient = &http.Client{}

// UpstashRequest sends a Redis command to Upstash via REST API
func UpstashRequest(command ...interface{}) (interface{}, error) {
	restURL := os.Getenv("UPSTASH_REDIS_REST_URL")
	token := os.Getenv("UPSTASH_REDIS_REST_TOKEN")

	if restURL == "" || token == "" {
		log.Fatal("UPSTASH_REDIS_REST_URL or UPSTASH_REDIS_REST_TOKEN not set")
	}

	// Marshal the Redis command (e.g., ["GET", "key"]) to JSON
	payload, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", restURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response (Upstash returns an array or string)
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// InitRedis tests the connection with PING
func InitRedis() {
	log.Println("Connecting to Upstash Redis (REST API)...")

	// Test with PING
	res, err := UpstashRequest("PING")
	if err != nil {
		log.Fatalf("Failed to connect to Upstash Redis: %v", err)
	}

	log.Printf("Connected to Upstash Redis. PING response: %v", res) // Should print "PONG"
}
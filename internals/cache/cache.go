package cache

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var cache *redis.Client

func Client () *redis.Client {
	return cache
}

func Connect(){
	ctx := context.Background()
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
		cache = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}else {
		redisOption , err := redis.ParseURL(redisURL)
		if err != nil{
			log.Fatalf("Failed to parse Redis URL: %v", err)
		}
		cache = redis.NewClient(redisOption)
	}
	cmd := cache.Ping(ctx)
	if cmd.Err() != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", cmd.Err()))
	}
	fmt.Println("Successfully Connected to Redis")
	
}
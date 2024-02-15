package config

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// ConnectToRedis creates and returns a new Redis client based on the provided configuration.
func ConnectToRedis(cfg *Config) (*redis.Client, error) {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHOST,
		Password: cfg.REDISPassword,
		DB:       0,
	})

	// Ping the Redis server to verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		// Return an error if the connection failed
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	// Connection successful, return the client
	return client, nil
}

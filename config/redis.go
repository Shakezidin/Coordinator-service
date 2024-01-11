package config

import "github.com/go-redis/redis/v8"

func ConnectToRedis(cfg *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHOST,
		Password: "NEW_PASSWORD",
		DB:       0,
	})
	return client
}

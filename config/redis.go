package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectRedis(cfg Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Successfully connected to Redis!")
	return rdb
}

func CacheExample(rdb *redis.Client) {
	// Menyimpan data ke Redis
	err := rdb.Set(ctx, "key", "Afdi Fauzul Bahar", 0).Err()
	if err != nil {
		log.Fatalf("Failed to set cache: %v", err)
	}

	// Mengambil data dari Redis
	val, err := rdb.Get(ctx, "key").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		log.Fatalf("Failed to get cache: %v", err)
	} else {
		fmt.Println("key:", val)
	}
}

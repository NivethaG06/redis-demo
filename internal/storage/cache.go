package storage

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	redisclient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisClient{client: redisclient}
}

func (r *RedisClient) Close() {
	r.client.Close()
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, serialized, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil
	} else if err != redis.Nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

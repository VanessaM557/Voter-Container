package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string, password string, db int) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{client: rdb}, nil
}

func (rc *RedisClient) Set(key string, value interface{}) error {
	return rc.client.Set(ctx, key, value, 0).Err()
}

func (rc *RedisClient) Get(key string) (string, error) {
	return rc.client.Get(ctx, key).Result()
}

func (rc *RedisClient) Del(key string) error {
	return rc.client.Del(ctx, key).Err()
}

func (rc *RedisClient) Keys(pattern string) ([]string, error) {
	return rc.client.Keys(ctx, pattern).Result()
}

package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *redisCache {
	return &redisCache{client: client}
}

func (c *redisCache) Get(key string) (string, error) {
	val, err := c.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (c *redisCache) Set(key string, value string, ttlSeconds int) error {
	return c.client.Set(context.Background(), key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

func (c *redisCache) Delete(key string) error {
	return c.client.Del(context.Background(), key).Err()
}

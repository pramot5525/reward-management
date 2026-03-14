package datasource

import (
	"github.com/pramot5525/reward-management/pkg/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg config.CacheConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}

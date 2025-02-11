package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crabocod/golang-test/internal/service"
	"time"

	"github.com/Crabocod/golang-test/internal/model"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewCacheService(host string, port int, db int, ttl time.Duration) service.CacheService {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		DB:   db,
	})

	return &redisCache{
		client: client,
		ttl:    ttl,
	}
}

func (s *redisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (s *redisCache) Set(ctx context.Context, key string, value *model.HashResponse) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value for cache: %w", err)
	}

	return s.client.Set(ctx, key, marshal, s.ttl).Err()
}

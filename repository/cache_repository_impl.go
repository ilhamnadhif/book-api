package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type CacheRepositoryImpl struct {
	Redis *redis.Client
	TTL   time.Duration
}

func NewCacheRepository(redis *redis.Client, ttl time.Duration) CacheRepository {
	return &CacheRepositoryImpl{
		Redis: redis,
		TTL:   ttl,
	}
}

func (repository *CacheRepositoryImpl) Get(ctx context.Context, key string, dest interface{}) error {
	res := repository.Redis.Get(ctx, key)
	if err := res.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		} else {
			return err
		}
	}
	bytes, err := res.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dest)
}

func (repository *CacheRepositoryImpl) Set(ctx context.Context, key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return repository.Redis.Set(ctx, key, bytes, repository.TTL).Err()
}

func (repository *CacheRepositoryImpl) Del(ctx context.Context, key string) error {
	return repository.Redis.Del(ctx, key).Err()
}

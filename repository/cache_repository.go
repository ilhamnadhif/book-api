package repository

import (
	"context"
)

type CacheRepository interface {
	Get(ctx context.Context, key string, dest interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	Del(ctx context.Context, key string) error
}

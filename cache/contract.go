package cache

import (
	"context"
	"time"
)

type Client interface {
	Keys(ctx context.Context, pattern string) ([]string, error)
	Get(ctx context.Context, key string, data interface{}) error
	Put(ctx context.Context, key string, data interface{}, expTime time.Duration) error
	Remove(ctx context.Context, key string) error
	RemoveAll(ctx context.Context, keys []string) error
	Close() error
}

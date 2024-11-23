package cache

import (
	"time"
)

type Client interface {
	Get(key string, data interface{}) error
	Put(key string, data interface{}, expTime time.Duration) error
	Remove(key string) error
	Close() error
}

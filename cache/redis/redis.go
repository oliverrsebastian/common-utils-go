package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type client struct {
	redisClient *redis.Client
}

func New(config *Config) *client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       0,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return &client{
		redisClient: redisClient,
	}
}

func (c *client) Close() error {
	return c.redisClient.Close()
}

func (c *client) Keys(ctx context.Context, pattern string) ([]string, error) {
	return c.redisClient.Keys(ctx, pattern).Result()
}

func (c *client) Get(ctx context.Context, key string, data interface{}) error {
	res := c.redisClient.Get(ctx, key)
	strData, err := res.Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(strData), &data)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Put(ctx context.Context, key string, data interface{}, expTime time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res := c.redisClient.Set(ctx, key, string(jsonData), expTime)
	return res.Err()
}

func (c *client) Remove(ctx context.Context, key string) error {
	res := c.redisClient.Del(ctx, key)
	return res.Err()
}

func (c *client) RemoveAll(ctx context.Context, keys []string) error {
	res := c.redisClient.Del(ctx, keys...)
	return res.Err()
}

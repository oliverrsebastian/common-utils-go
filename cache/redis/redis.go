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

func (cm *client) Close() error {
	return cm.redisClient.Close()
}

func (cm *client) Get(key string, data interface{}) error {
	res := cm.redisClient.Get(context.Background(), key)
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

func (cm *client) Put(key string, data interface{}, expTime time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res := cm.redisClient.Set(context.Background(), key, string(jsonData), expTime)
	return res.Err()
}

func (cm *client) Remove(key string) error {
	res := cm.redisClient.Del(context.Background(), key)
	return res.Err()
}

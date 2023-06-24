package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Client  *redis.Client
	Context context.Context
}

func NewRedisService() *RedisService {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisService{
		Client:  client,
		Context: context.Background(),
	}
}

func (r *RedisService) AddKey(key string, value float64) error {
	_, err := r.Client.Set(r.Context, key, value, 60*time.Minute).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisService) ReadKey(key string) (*string, error) {
	value, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		return nil, err
	}

	return &value, nil
}

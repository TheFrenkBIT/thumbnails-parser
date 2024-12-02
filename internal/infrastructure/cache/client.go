package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	*redis.Client
}

func New(client *redis.Client) *Redis {
	return &Redis{
		client,
	}
}

func (r *Redis) SetValue(ctx context.Context, key string, value interface{}) error {
	err := r.Set(ctx, key, value, time.Minute*3).Err()

	return err
}
func (r *Redis) GetValue(ctx context.Context, key string) (interface{}, error) {
	value, err := r.Get(ctx, key).Bytes()

	return value, err
}

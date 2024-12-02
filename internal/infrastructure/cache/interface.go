package cache

import "context"

type Interface interface {
	SetValue(ctx context.Context, key string, value interface{}) error
	GetValue(ctx context.Context, key string) (interface{}, error)
}

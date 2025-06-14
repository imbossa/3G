package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Option defines a functional option for RedisOption.
type Option func(*RedisOption)

// RedisOption holds Redis client configuration.
type RedisOption struct {
	PoolSize        int
	PoolTimeout     time.Duration
	ConnMaxIdleTime time.Duration
}

// Redis wraps the redis.Client.
type Redis struct {
	URL    string
	Option RedisOption
	Client *redis.Client
}

// New creates a new Redis client instance using github.com/redis/go-redis/v9's ParseURL.
// Only pool parameters can be set via Option; all connection info is from the URL.
func New(urlStr string, opts ...Option) (*Redis, error) {
	option := RedisOption{
		PoolSize:        10,
		PoolTimeout:     30 * time.Second,
		ConnMaxIdleTime: 5 * time.Minute,
	}

	for _, opt := range opts {
		opt(&option)
	}

	redisOpt, err := redis.ParseURL(urlStr)
	if err != nil {
		return nil, fmt.Errorf("redis - New - ParseURL: %w", err)
	}
	redisOpt.PoolSize = option.PoolSize
	redisOpt.PoolTimeout = option.PoolTimeout
	redisOpt.ConnMaxIdleTime = option.ConnMaxIdleTime

	client := redis.NewClient(redisOpt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis - New - client.Ping: %w", err)
	}

	return &Redis{
		URL:    urlStr,
		Option: option,
		Client: client,
	}, nil
}

// Close closes the Redis client connection.
func (r *Redis) Close() error {
	if r.Client == nil {
		return nil
	}
	return r.Client.Close()
}

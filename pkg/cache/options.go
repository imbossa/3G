package redis

import "time"

// WithPoolSize sets the connection pool size.
func WithPoolSize(size int) Option {
	return func(opt *RedisOption) {
		opt.PoolSize = size
	}
}

// WithPoolTimeout sets the pool timeout duration.
func WithPoolTimeout(timeout time.Duration) Option {
	return func(opt *RedisOption) {
		opt.PoolTimeout = timeout
	}
}

// WithConnMaxIdleTime sets the maximum idle time for a connection.
func WithConnMaxIdleTime(d time.Duration) Option {
	return func(opt *RedisOption) {
		opt.ConnMaxIdleTime = d
	}
}
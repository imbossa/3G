package db

// Option type for functional options.
type Option func(*DBOption)

// MaxIdleConns sets the maximum number of idle connections.
func MaxIdleConns(n int) Option {
	return func(opt *DBOption) {
		opt.MaxIdleConns = n
	}
}

// MaxOpenConns sets the maximum number of open connections.
func MaxOpenConns(n int) Option {
	return func(opt *DBOption) {
		opt.MaxOpenConns = n
	}
}

// Config sets a custom config string.
func Config(cfg string) Option {
	return func(opt *DBOption) {
		opt.CONFIG = cfg
	}
}

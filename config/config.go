package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type (
	// Config -.
	Config struct {
		App     App
		HTTP    HTTP
		Log     Log
		DB      DB
		REDIS   REDIS
		GRPC    GRPC
		RMQ     RMQ
		Swagger Swagger
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}

	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	// DB -.
	DB struct {
		URL          string `env:"DB_URL,required"`
		MaxIdleConns int    `env:"DB_OPTION_MAX_IDLE_CONNS" envDefault:"5"`
		MaxOpenConns int    `env:"DB_OPTION_MAX_OPEN_CONNS" envDefault:"10"`
		CONFIG       string `env:"DB_OPTION_CONFIG"`
	}

	// REDIS -.
	REDIS struct {
		URL         string `env:"REDIS_URL,required"`
		POOL        int `env:"REDIS_POOl_SIZE" envDefault:"10"`
		TIMEOUT     int `env:"REDIS_POOl_TIMEOUT" envDefault:"10"`
		MaxIdleTime int `env:"REDIS_CONN_MAX_IDLE_TIME" envDefault:"120"`
	}

	// GRPC -.
	GRPC struct {
		Port string `env:"GRPC_PORT,required"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env:"RMQ_RPC_SERVER,required"`
		ClientExchange string `env:"RMQ_RPC_CLIENT,required"`
		URL            string `env:"RMQ_URL,required"`
	}

	// Swagger -.
	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

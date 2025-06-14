// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/imbossa/3G/config"
	amqprpc "github.com/imbossa/3G/internal/controller/amqp_rpc"
	"github.com/imbossa/3G/internal/controller/grpc"
	"github.com/imbossa/3G/internal/controller/http"
	"github.com/imbossa/3G/internal/repo/persistent"
	"github.com/imbossa/3G/internal/repo/webapi"
	"github.com/imbossa/3G/internal/usecase/translation"
	cache "github.com/imbossa/3G/pkg/cache"
	"github.com/imbossa/3G/pkg/db"
	"github.com/imbossa/3G/pkg/grpcserver"
	"github.com/imbossa/3G/pkg/httpserver"
	"github.com/imbossa/3G/pkg/logger"
	"github.com/imbossa/3G/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	database, err := db.New(cfg.DB.URL, db.MaxIdleConns(cfg.DB.MaxIdleConns), db.MaxOpenConns(cfg.DB.MaxOpenConns), db.Config(cfg.DB.CONFIG))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - db.New: %w", err))
	}
	defer database.Close()

	// Use-Case
	translationUseCase := translation.New(
		persistent.New(database),
		webapi.New(),
	)

	// redis
	redis, err := cache.New(cfg.REDIS.URL,
		cache.WithPoolSize(cfg.REDIS.POOL),
		cache.WithPoolTimeout(time.Duration(cfg.REDIS.TIMEOUT)*time.Second),
		cache.WithConnMaxIdleTime(time.Duration(cfg.REDIS.MaxIdleTime)*time.Second),
	)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - cache.New: %w", err))
	}
	defer redis.Close()

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(translationUseCase, l)
	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// gRPC Server
	grpcServer := grpcserver.New(grpcserver.Port(cfg.GRPC.Port))
	grpc.NewRouter(grpcServer.App, translationUseCase, l)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	http.NewRouter(httpServer.Engine, cfg, translationUseCase, l, redis)

	// Start servers
	rmqServer.Start()
	grpcServer.Start()
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-grpcServer.Notify():
		l.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	httpServerErr := httpServer.Shutdown()
	if httpServerErr != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	grpcServerErr := grpcServer.Shutdown()
	if grpcServerErr != nil {
		l.Error(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
	}

	rmqServerErr := rmqServer.Shutdown()
	if rmqServerErr != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}

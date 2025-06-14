// Package v1 implements routing paths. Each services in own file.
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/imbossa/3G/config"
	_ "github.com/imbossa/3G/docs" // Swagger docs.
	example "github.com/imbossa/3G/internal/controller/http/example"
	"github.com/imbossa/3G/internal/controller/http/middleware"
	"github.com/imbossa/3G/internal/usecase"
	redis "github.com/imbossa/3G/pkg/cache"
	"github.com/imbossa/3G/pkg/logger"
	// Swagger for Gin
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter -.
// Swagger spec:
// @title Go Clean Template API
// @description Using a translation service as an example
// @version 1.0
// @host localhost:8080
// @BasePath /example
func NewRouter(router *gin.Engine, cfg *config.Config, t usecase.Translation, l logger.Interface, c *redis.Redis) {
	// Options: Logger & Recovery
	router.Use(middleware.Logger(l))
	router.Use(middleware.Recovery(l))

	// Swagger
	if cfg.Swagger.Enabled {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Routers
	apiExampleGroup := router.Group("/example")
	{
		example.NewTranslationRoutes(apiExampleGroup, t, l, c)
	}
}

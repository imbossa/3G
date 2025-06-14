package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/imbossa/3G/internal/usecase"
	redis "github.com/imbossa/3G/pkg/cache"
	"github.com/imbossa/3G/pkg/logger"
)

// NewTranslationRoutes - Gin 版
func NewTranslationRoutes(apiExampleGroup *gin.RouterGroup, t usecase.Translation, l logger.Interface,c *redis.Redis) {
	r := &Example{t: t, l: l, c: c, v: validator.New(validator.WithRequiredStructEnabled())}

	translationGroup := apiExampleGroup.Group("/translation")
	{
		translationGroup.GET("/history", r.history)
		translationGroup.POST("/do-translate", r.doTranslate)
	}
}

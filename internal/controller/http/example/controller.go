package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/imbossa/3G/internal/usecase"
	"github.com/imbossa/3G/pkg/cache"
	"github.com/imbossa/3G/pkg/logger"
)

// Example -.
type Example struct {
	t usecase.Translation
	l logger.Interface
	v *validator.Validate
	c *redis.Redis
}

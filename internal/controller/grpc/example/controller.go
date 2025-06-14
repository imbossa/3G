package grpc

import (
	"github.com/go-playground/validator/v10"
	protoexample "github.com/imbossa/3G/docs/proto/example"
	"github.com/imbossa/3G/internal/usecase"
	"github.com/imbossa/3G/pkg/logger"
)

// Example -.
type Example struct {
	protoexample.TranslationServer

	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}

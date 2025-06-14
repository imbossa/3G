package grpc

import (
	"github.com/go-playground/validator/v10"
	protoexample "github.com/imbossa/3G/docs/proto/example"
	"github.com/imbossa/3G/internal/usecase"
	"github.com/imbossa/3G/pkg/logger"
	pbgrpc "google.golang.org/grpc"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(app *pbgrpc.Server, t usecase.Translation, l logger.Interface) {
	r := &Example{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}
	{
		protoexample.RegisterTranslationServer(app, r)
	}
}

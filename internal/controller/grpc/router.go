package grpc

import (
	example "github.com/imbossa/3G/internal/controller/grpc/example"
	"github.com/imbossa/3G/internal/usecase"
	"github.com/imbossa/3G/pkg/logger"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewRouter -.
func NewRouter(app *pbgrpc.Server, t usecase.Translation, l logger.Interface) {
	{
		example.NewTranslationRoutes(app, t, l)
	}

	reflection.Register(app)
}

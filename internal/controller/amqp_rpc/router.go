package amqp_rpc

import (
	example "github.com/imbossa/3G/internal/controller/amqp_rpc/example"
	"github.com/imbossa/3G/internal/usecase"
	"github.com/imbossa/3G/pkg/logger"
	"github.com/imbossa/3G/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation, l logger.Interface) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		example.NewTranslationRoutes(routes, t, l)
	}

	return routes
}

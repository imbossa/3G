package amqp_rpc

import (
	"github.com/go-playground/validator/v10"
	"github.com/imbossa/3G/internal/usecase"
	"github.com/imbossa/3G/pkg/logger"
	"github.com/imbossa/3G/pkg/rabbitmq/rmq_rpc/server"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(routes map[string]server.CallHandler, t usecase.Translation, l logger.Interface) {
	r := &Example{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}
	{
		routes["example.getHistory"] = r.getHistory()
	}
}

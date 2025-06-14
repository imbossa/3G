package amqp_rpc

import (
	"context"
	"fmt"

	"github.com/imbossa/3G/pkg/rabbitmq/rmq_rpc/server"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *Example) getHistory() server.CallHandler {
	return func(_ *amqp.Delivery) (interface{}, error) {
		translationHistory, err := r.t.History(context.Background())
		if err != nil {
			r.l.Error(err, "amqp_rpc - V1 - getHistory")

			return nil, fmt.Errorf("amqp_rpc - V1 - getHistory: %w", err)
		}

		return translationHistory, nil
	}
}

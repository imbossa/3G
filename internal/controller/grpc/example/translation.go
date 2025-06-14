package grpc

import (
	"context"
	"fmt"

	protoexample "github.com/imbossa/3G/docs/proto/example"
	"github.com/imbossa/3G/internal/controller/grpc/example/response"
)

func (r *Example) GetHistory(ctx context.Context, _ *protoexample.GetHistoryRequest) (*protoexample.GetHistoryResponse, error) {
	translationHistory, err := r.t.History(ctx)
	if err != nil {
		r.l.Error(err, "grpc - example - GetHistory")

		return nil, fmt.Errorf("grpc - example - GetHistory: %w", err)
	}

	return response.NewTranslationHistory(translationHistory), nil
}

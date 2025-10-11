package order

import (
	"context"
	"order/internal/client/converter"
	"order/pkg/api"
)

func (a *Server) HandleCancelOrder(ctx context.Context, req api.HandleCancelOrderParams) (api.HandleCancelOrderRes, error) {
	err := a.serv.CancelOrder(ctx, converter.UUIDToString(req.OrderUUID))
	if err != nil {
		return nil, err
	}
	return &api.HandleCancelOrderNoContent{}, nil
}

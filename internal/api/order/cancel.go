package order

import (
	"context"
	"order/internal/client/converter"
	"order/internal/entity"
	"order/pkg/api"
)

func (a *Server) HandleCancelOrder(ctx context.Context, req api.HandleCancelOrderParams) (api.HandleCancelOrderRes, error) {
	err := a.serv.CancelOrder(ctx, converter.UUIDToString(req.OrderUUID))
	if err != nil {
		return nil, entity.ErrInternalError
	}
	return &api.HandleCancelOrderNoContent{}, nil
}

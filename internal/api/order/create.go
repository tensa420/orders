package order

import (
	"context"
	"order/internal/client/converter"
	"order/internal/entity"
	"order/pkg/api"
)

func (a *Server) HandleCreateOrder(ctx context.Context, req *api.CreateOrderRequest) (api.HandleCreateOrderRes, error) {
	orderUUID, total, err := a.serv.CreateOrder(ctx, req.UserUUID.String(), req.PartUuids)
	if err != nil {
		return nil, entity.ErrInternalError
	}
	return &api.CreateOrderResponse{
		OrderUUID:  converter.StringToUUID(orderUUID),
		TotalPrice: total,
	}, nil
}

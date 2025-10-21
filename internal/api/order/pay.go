package order

import (
	"context"
	"order/internal/client/converter"
	"order/internal/entity"
	"order/pkg/api"
)

func (a *Server) HandlePayOrder(ctx context.Context, req *api.PayOrderRequest, params api.HandlePayOrderParams) (api.HandlePayOrderRes, error) {
	transuuid, err := a.serv.PayOrder(ctx, params.OrderUUID.String(), req.PaymentMethod)
	if err != nil {
		return nil, entity.ErrInternalError
	}
	return &api.PayOrderResponse{TransactionUUID: converter.StringToUUID(transuuid)}, nil
}

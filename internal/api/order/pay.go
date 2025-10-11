package order

import (
	"context"
	"order/internal/client/converter"
	"order/pkg/api"
)

func (a *Server) HandlePayOrder(ctx context.Context, req *api.PayOrderRequest, params api.HandlePayOrderParams) (api.HandlePayOrderRes, error) {
	transuuid, err := a.serv.PayOrder(ctx, params.OrderUUID.String(), req.PaymentMethod)
	if err != nil {
		return nil, err
	}
	return &api.PayOrderResponse{TransactionUUID: converter.StringToUUID(transuuid)}, nil
}

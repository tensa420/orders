package order

import (
	"context"
	ap "order/api"
	"order/internal/client/converter"
)

func (a *api) PayOrder(ctx context.Context, req ap.PayOrderRequest, params ap.HandlePayOrderParams) (ap.HandlePayOrderRes, error) {
	transuuid, err := a.serv.PayOrder(ctx, params.OrderUUID.String(), req.PaymentMethod)
	if err != nil {
		return nil, err
	}
	return &ap.PayOrderResponse{TransactionUUID: converter.StringToUUID(transuuid)}, nil
}

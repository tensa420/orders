package order

import (
	"context"
	ap "order/api"
	"order/internal/client/converter"
)

func (a *Api) HandleCancelOrder(ctx context.Context, req ap.HandleCancelOrderParams) (ap.HandleCancelOrderRes, error) {
	err := a.serv.CancelOrder(ctx, converter.UUIDToString(req.OrderUUID))
	if err != nil {
		return nil, err
	}
	return &ap.HandleCancelOrderNoContent{}, nil
}

package api

import (
	"context"
	"order/api"
)

type OrderServer interface {
	HandleGetOrder(ctx context.Context, params api.HandleGetOrderParams) (api.HandleGetOrderRes, error)
	HandleCancelOrder(ctx context.Context, params api.HandleCancelOrderParams) (api.HandleCancelOrderRes, error)
	HandlePayOrder(ctx context.Context, req api.PayOrderRequest, params api.HandlePayOrderParams) (api.HandlePayOrderRes, error)
	HandleCreateOrder(ctx context.Context, req api.CreateOrderRequest) (api.HandleCreateOrderRes, error)
}

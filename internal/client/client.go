package client

import (
	"context"
	"order/api"
)

type Client interface {
	GetOrder(ctx context.Context, params api.HandleGetOrderParams) (*api.HandleGetOrderRes, error)
	CancelOrder(ctx context.Context, params api.HandleCancelOrderParams) (api.HandleCancelOrderRes, error)
	PayOrder(ctx context.Context, req api.PayOrderRequest, params api.HandlePayOrderParams) (api.HandlePayOrderRes, error)
	CreateOrder(ctx context.Context, req api.CreateOrderRequest) (api.HandleCreateOrderRes, error)
}

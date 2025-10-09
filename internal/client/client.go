package client

import (
	"context"
	"order/internal/repository/model"
)

type Client interface {
	GetOrder(ctx context.Context, orderUUID string) (*model.GetOrderResponse, error)
	CancelOrder(ctx context.Context, orderUUID string) error
	PayOrder(ctx context.Context, orderUUID string, paymentMethod string) error
	CreateOrder(ctx context.Context, userUUID string, partsUUID []string) (string, error)
}

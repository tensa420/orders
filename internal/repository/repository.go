package repository

import (
	"context"
	"order/internal/repository/model"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, userUUID string, parts []model.Part, total float64) (string, error)
	PayOrder(ctx context.Context, transUUID string, orderUUID string) error
	GetOrder(ctx context.Context, orderUUID string) (*model.GetOrderResponse, error)
	CancelOrder(ctx context.Context, orderUUID string) error
}

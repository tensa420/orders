package service

import (
	"context"
	"order/internal/entity"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userUUID string, partUUIDS []string) (string, float64, error)
	PayOrder(ctx context.Context, orderUUID string, paymentMethod string) (string, error)
	GetOrder(ctx context.Context, orderUUID string) (*entity.GetOrderResponse, error)
	CancelOrder(ctx context.Context, orderUUID string) error
}

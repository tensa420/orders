package repository

import (
	"context"
	"order/internal/entity"
)

type OrderRepository interface {
	GetOrder(ctx context.Context, orderUUID string) (*entity.Order, error)
	CreateOrder(ctx context.Context, order entity.Order) error
	PayOrder(ctx context.Context, info entity.PaymentInfo) error
	CancelOrder(ctx context.Context, orderUUID string) error
}

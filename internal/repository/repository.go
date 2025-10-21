package repository

import (
	"context"
	"order/internal/entity"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, userUUID string, partUUIDS []string, total float64) (string, error)
	PayOrder(ctx context.Context, transUUID string, orderUUID string, paymentMethod entity.PaymentMethod) error
	GetOrder(ctx context.Context, orderUUID string) (*entity.Order, error)
	CancelOrder(ctx context.Context, orderUUID string) error
}

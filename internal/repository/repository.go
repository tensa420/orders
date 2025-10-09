package repository

import (
	"context"
	repoModel "order/internal/repository/model"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, userUUID string, parts []repoModel.Part, total float64) (string, error)
	PayOrder(ctx context.Context, transUUID string, orderUUID string, paymentMethod repoModel.PaymentMethod) error
	GetOrder(ctx context.Context, orderUUID string) (*repoModel.GetOrderResponse, error)
	CancelOrder(ctx context.Context, orderUUID string) error
}

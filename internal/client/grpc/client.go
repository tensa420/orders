package grpc

import (
	"context"
	"order/internal/entity"
)

type InventoryClient interface {
	ListParts(ctx context.Context, partsUUID []string) ([]*entity.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUUID string, userUUID string, PaymentMethod string) (string, error)
}

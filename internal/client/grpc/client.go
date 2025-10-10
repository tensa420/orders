package grpc

import (
	"context"
	repoModel "order/internal/repository/model"
)

type InventoryClient interface {
	ListParts(ctx context.Context, partsUUID []string) ([]*repoModel.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUUID string, userUUID string, PaymentMethod string) (string, error)
}

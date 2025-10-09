package grpc

import (
	"context"
	"order/internal/repository/model"
)

type InventoryClient interface {
	ListParts(ctx context.Context, partUUIDS []string) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, userUUID, orderUUID, PaymentMethod string) (string, error)
}

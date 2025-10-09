package order

import (
	"context"
	mod "order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *repository) CancelOrder(ctx context.Context, orderUUID string) error {

	ord, ok := r.orders[orderUUID]
	if !ok {
		return status.Error(codes.NotFound, "order not found")
	}

	if ord.Status == mod.Status(0) {
		return status.Error(codes.Internal, "Order already paid")
	}

	ord.Status = mod.Status(2)
	return nil
}

package order

import (
	"context"
	repoModel "order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *Repository) CancelOrder(ctx context.Context, orderUUID string) error {

	ord, ok := r.orders[orderUUID]
	if !ok {
		return status.Error(codes.NotFound, "order not found")
	}

	if ord.Status == repoModel.StatusPaid {
		return status.Error(codes.Internal, "Order already paid")
	}

	ord.Status = repoModel.StatusCancelled
	return nil
}

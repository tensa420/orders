package order

import (
	"context"
	"order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *repository) PayOrder(ctx context.Context, transUUID string, orderUUID string) error {
	r.mu.Lock()
	ord, ok := r.orders[orderUUID]
	r.mu.Unlock()

	if !ok {
		return status.Error(codes.NotFound, "order not found")
	}

	r.mu.Lock()
	ord.Status = model.Status(0)
	ord.TransactionUUID = &transUUID
	r.mu.Unlock()

	return nil
}

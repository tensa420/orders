package order

import (
	"context"
	repoModel "order/internal/repository/model"
)

func (r *OrderRepository) CancelOrder(ctx context.Context, orderUUID string) error {
	r.mu.RLock()
	ord, ok := r.orders[orderUUID]
	r.mu.RUnlock()

	if !ok {
		return repoModel.ErrOrderNotFound
	}

	if ord.Status == repoModel.StatusPaid {
		return repoModel.ErrInternalError
	}

	ord.Status = repoModel.StatusCancelled
	return nil
}

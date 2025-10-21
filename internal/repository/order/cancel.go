package order

import (
	"context"
	"order/internal/entity"
	repoModel "order/internal/repository/model"
)

func (r *OrderRepository) CancelOrder(ctx context.Context, orderUUID string) error {
	r.mu.RLock()
	ord, ok := r.orders[orderUUID]
	r.mu.RUnlock()

	if !ok {
		return entity.ErrOrderNotFound
	}

	if ord.Status == repoModel.StatusPaid {
		return entity.ErrInternalError
	}

	ord.Status = repoModel.StatusCancelled
	return entity.ErrSuccessCancel
}

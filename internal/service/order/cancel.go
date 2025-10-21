package order

import (
	"context"
	"order/internal/entity"
)

func (s *OrderService) CancelOrder(ctx context.Context, orderUUID string) error {
	err := s.repo.CancelOrder(ctx, orderUUID)
	if err != nil {
		return entity.ErrInternalError
	}
	return entity.ErrSuccessCancel
}

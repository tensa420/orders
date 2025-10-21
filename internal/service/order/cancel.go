package order

import (
	"context"
)

func (s *OrderService) CancelOrder(ctx context.Context, orderUUID string) error {
	err := s.repo.CancelOrder(ctx, orderUUID)
	if err != nil {
		return err
	}
	return nil
}

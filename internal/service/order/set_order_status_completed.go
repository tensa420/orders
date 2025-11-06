package order

import "context"

func (s *OrderService) SetOrderStatusCompleted(ctx context.Context, orderUUID string) error {
	err := s.repo.SetOrderStatusCompleted(ctx, orderUUID)
	if err != nil {
		return err
	}
	return nil
}

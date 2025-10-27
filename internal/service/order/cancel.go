package order

import (
	"context"
	"order/internal/entity"

	"github.com/go-faster/errors"
)

func (s *OrderService) CancelOrder(ctx context.Context, orderUUID string) error {
	order, err := s.repo.GetOrder(ctx, orderUUID)
	if err != nil {
		return entity.ErrOrderNotFound
	}

	if order.Status == "PAID" {
		return errors.Wrap(err, "order already paid")
	}

	err = s.repo.CancelOrder(ctx, orderUUID)
	if err != nil {
		return err
	}
	return nil
}

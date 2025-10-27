package order

import (
	"context"
	"order/internal/entity"

	"github.com/google/uuid"
)

func (s *OrderService) CreateOrder(ctx context.Context, userUUID string, partsUUID []string) (string, float64, error) {
	total := 0.0

	parts, err := s.invClient.ListParts(ctx, partsUUID)
	if err != nil {
		return "", 0, err
	}

	for _, part := range parts {
		total += part.Price
	}

	orderUUID := uuid.New().String()

	order := entity.Order{
		OrderUUID:  orderUUID,
		UserUUID:   userUUID,
		PartsUUID:  partsUUID,
		TotalPrice: total,
		Status:     entity.StatusPendingPayment,
	}

	err = s.repo.CreateOrder(ctx, order)
	if err != nil {
		return "", 0, err
	}

	return orderUUID, total, nil
}

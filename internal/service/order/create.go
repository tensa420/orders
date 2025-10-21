package order

import (
	"context"
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
	orderUUID, err := s.repo.CreateOrder(ctx, userUUID, partsUUID, total)
	if err != nil {
		return "", 0, err
	}
	return orderUUID, total, nil
}

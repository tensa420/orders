package order

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CreateOrder(ctx context.Context, userUUID string, partsUUID []string) (string, float64, error) {
	total := 0.0
	parts, err := s.InventoryClient.ListParts(ctx, partsUUID)
	if err != nil {
		return "", 0, status.Error(codes.Internal, err.Error())
	}
	for _, part := range parts {
		total += part.Price
	}
	orderUUID, err := s.OrderRepository.CreateOrder(ctx, userUUID, parts, total)
	if err != nil {
		return "", 0, status.Error(codes.Internal, err.Error())
	}
	return orderUUID, total, nil
}

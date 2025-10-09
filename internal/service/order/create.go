package order

import (
	"context"
	"order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CreateOrder(ctx context.Context, order model.Order) (string, float64, error) {
	total := 0.0
	parts, err := s.InventoryClient.ListParts(ctx, order.PartsUUID)
	if err != nil {
		return "", 0, status.Error(codes.Internal, err.Error())
	}
	for _, part := range parts {
		total += part.Price
	}
	ord, err := s.OrderRepository.CreateOrder(ctx, order.UserUUID, parts, total)
	if err != nil {
		return "", 0, status.Error(codes.Internal, err.Error())
	}
	return ord, total, nil
}

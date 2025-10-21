package order

import (
	"context"
	"order/internal/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) GetOrder(ctx context.Context, orderUUID string) (*entity.Order, error) {
	ord, err := s.repo.GetOrder(ctx, orderUUID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &entity.Order{
		OrderUUID:       ord.OrderUUID,
		UserUUID:        ord.UserUUID,
		TransactionUUID: ord.TransactionUUID,
		PaymentMethod:   ord.PaymentMethod,
		PartsUUID:       ord.PartsUUID,
		Status:          ord.Status,
		TotalPrice:      ord.TotalPrice,
	}, nil
}

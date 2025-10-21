package order

import (
	"context"
	"order/internal/entity"
)

func (s *OrderService) GetOrder(ctx context.Context, orderUUID string) (*entity.Order, error) {
	ord, err := s.repo.GetOrder(ctx, orderUUID)
	if err != nil {
		return nil, err
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

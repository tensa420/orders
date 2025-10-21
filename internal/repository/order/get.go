package order

import (
	"context"
	"order/internal/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *OrderRepository) GetOrder(ctx context.Context, orderUUID string) (*entity.Order, error) {
	r.mu.RLock()
	ord, ok := r.orders[orderUUID]
	r.mu.RUnlock()

	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &entity.Order{
		OrderUUID:       orderUUID,
		UserUUID:        ord.UserUUID,
		PartsUUID:       ord.PartsUUID,
		TotalPrice:      ord.TotalPrice,
		TransactionUUID: ord.TransactionUUID,
		PaymentMethod:   ord.PaymentMethod,
		Status:          ord.Status,
	}, nil
}

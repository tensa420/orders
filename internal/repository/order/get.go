package order

import (
	"context"
	"order/internal/entity"
	repoModel "order/internal/repository/model"
)

func (r *OrderRepository) GetOrder(ctx context.Context, orderUUID string) (*entity.Order, error) {
	r.mu.RLock()
	ord, ok := r.orders[orderUUID]
	r.mu.RUnlock()

	if !ok {
		return nil, repoModel.ErrOrderNotFound
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

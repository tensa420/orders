package order

import (
	"context"
	repoModel "order/internal/repository/model"

	"github.com/google/uuid"
)

func (r *repository) CreateOrder(ctx context.Context, userUUID string, partUUIDS []string, total float64) (string, error) {
	OrderUUID := uuid.New()

	order := repoModel.Order{
		OrderUUID:  OrderUUID.String(),
		UserUUID:   userUUID,
		PartsUUID:  partUUIDS,
		TotalPrice: total,
		Status:     repoModel.StatusPendingPayment,
	}

	r.mu.Lock()
	r.orders[order.OrderUUID] = order
	r.mu.Unlock()

	return OrderUUID.String(), nil
}

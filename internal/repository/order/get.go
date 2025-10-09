package order

import (
	"context"
	repoModel "order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *repository) HandleGetOrder(ctx context.Context, orderUUID string) (*repoModel.GetOrderResponse, error) {
	r.mu.Lock()
	ord, ok := r.orders[orderUUID]
	r.mu.Unlock()

	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &repoModel.GetOrderResponse{
		OrderUUID:       orderUUID,
		UserUUID:        ord.UserUUID,
		PartUuids:       ord.PartsUUID,
		TotalPrice:      ord.TotalPrice,
		TransactionUUID: ord.TransactionUUID,
		PaymentMethod:   ord.PaymentMethod,
		Status:          ord.Status,
	}, nil
}

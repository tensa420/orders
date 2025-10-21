package order

import (
	"context"
	"order/internal/client/converter"
	"order/internal/entity"
	repoModel "order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *OrderRepository) PayOrder(ctx context.Context, transUUID string, orderUUID string, paymentMethod entity.PaymentMethod) error {
	r.mu.RLock()
	ord, ok := r.orders[orderUUID]
	r.mu.RUnlock()

	if !ok {
		return status.Error(codes.NotFound, "order not found")
	}

	finalPaymentMethod := converter.ConvertPaymentMethodToString(paymentMethod)

	r.mu.Lock()
	ord.Status = repoModel.StatusPaid
	ord.TransactionUUID = &transUUID
	ord.PaymentMethod = &finalPaymentMethod
	r.mu.Unlock()

	return nil
}

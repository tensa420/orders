package order

import (
	"context"
	"order/internal/client/converter"
	"order/internal/entity"
	repoModel "order/internal/repository/model"
)

func (r *OrderRepository) PayOrder(ctx context.Context, transUUID string, orderUUID string, paymentMethod entity.PaymentMethod) error {
	r.mu.RLock()
	ord, ok := r.orders[orderUUID]
	r.mu.RUnlock()

	if !ok {
		return repoModel.ErrOrderNotFound
	}

	finalPaymentMethod := converter.ConvertPaymentMethodToString(paymentMethod)

	r.mu.Lock()
	ord.Status = repoModel.StatusPaid
	ord.TransactionUUID = &transUUID
	ord.PaymentMethod = &finalPaymentMethod
	r.mu.Unlock()

	return nil
}

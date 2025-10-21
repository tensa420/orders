package order

import (
	"context"
	"order/internal/client/converter"
)

func (s *OrderService) PayOrder(ctx context.Context, orderUUID string, paymentMethod string) (string, error) {
	convertedPaymentMethod := converter.ConvertPaymentMethod(paymentMethod)

	ord, err := s.repo.GetOrder(ctx, orderUUID)
	if err != nil {
		return "", err
	}

	transactionUUID, err := s.paymClient.PayOrder(ctx, orderUUID, ord.UserUUID, paymentMethod)
	if err != nil {
		return "", err
	}

	err = s.repo.PayOrder(ctx, transactionUUID, orderUUID, convertedPaymentMethod)
	if err != nil {
		return "", err
	}

	return transactionUUID, nil
}

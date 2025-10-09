package payment

import (
	"context"
	"order/internal/client/converter"
	v1 "order/pkg/payment/payment"
)

func (c *client) PayOrder(ctx context.Context, userUUID, orderUUID, PaymentMethod string) (string, error) {
	req, err := c.generatedClient.PayOrder(ctx, &v1.PayOrderRequest{
		Order: &v1.OrderRequest{
			UserUuid:      userUUID,
			OrderUuid:     orderUUID,
			PaymentMethod: converter.PaymentMethodToEnum(PaymentMethod),
		},
	})
	if err != nil {
		return "", err
	}
	return req.TransactionUuid, nil
}

package payment

import (
	"context"
	"order/internal/client/converter"
	v1 "order/pkg/payment"
)

func (c *Client) PayOrder(ctx context.Context, orderUUID string, userUUID string, PaymentMethod string) (string, error) {
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

package payment

import (
	v1 "order/pkg/payment/payment"
)

type client struct {
	generatedClient v1.PaymentClient
}

func New(generatedClient v1.PaymentClient) *client {
	return &client{generatedClient: generatedClient}
}

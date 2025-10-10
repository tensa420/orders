package payment

import (
	v1 "order/pkg/payment"
)

type Client struct {
	generatedClient v1.PaymentClient
}

func New(generatedClient v1.PaymentClient) *Client {
	return &Client{generatedClient: generatedClient}
}

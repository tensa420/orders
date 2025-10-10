package order

import (
	"order/internal/api"
	"order/internal/service"
	"order/pkg/inventory"
	"order/pkg/payment"
)

var _ api.Client = (*Api)(nil)

type Api struct {
	invent  inventory.InventoryServiceClient
	payment payment.PaymentClient
	serv    service.OrderService
}

func NewAPI(inv inventory.InventoryServiceClient, serv service.OrderService, paym payment.PaymentClient) *Api {

	return &Api{
		serv:    serv,
		invent:  inv,
		payment: paym,
	}
}

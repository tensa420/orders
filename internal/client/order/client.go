package order

import (
	"order/internal/service"
	inv "order/pkg/inventory/inventory"
	paym "order/pkg/payment/payment"
)

type api struct {
	invent  inv.UnimplementedInventoryServiceServer
	payment paym.UnimplementedPaymentServer
	serv    service.OrderService
}

func NewAPI(inv inv.UnimplementedInventoryServiceServer, serv service.OrderService, paym paym.UnimplementedPaymentServer) *api {
	return &api{
		serv:    serv,
		invent:  inv,
		payment: paym,
	}
}

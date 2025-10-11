package order

import (
	"order/internal/api"
	"order/internal/service"
	apii "order/pkg/api"
	"order/pkg/inventory"
	"order/pkg/payment"
)

var _ api.OrderServer = (*Server)(nil)

type Server struct {
	invent inventory.InventoryServiceClient
	apii.UnimplementedHandler
	payment payment.PaymentClient
	serv    service.OrderService
}

func NewOrderServer(serv service.OrderService) *Server {
	return &Server{
		serv: serv,
	}
}

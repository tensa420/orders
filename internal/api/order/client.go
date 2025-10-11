package order

import (
	apii "order/api"
	"order/internal/api"
	"order/internal/service"
	"order/pkg/inventory"
	"order/pkg/payment"
)

var _ api.OrderServer = (*Server)(nil)

type Server struct {
	invent  inventory.InventoryServiceClient
	payment payment.PaymentClient
	serv    service.OrderService
}

func NewAPI(serv service.OrderService) *Server {

	return &Server{
		serv: serv,
	}
}

type OrderHandler struct {
	serv Server
	apii.UnimplementedHandler
	paym payment.PaymentClient
	inv  inventory.InventoryServiceClient
}

func NewOrderHandler(serv *Server) *OrderHandler {
	return &OrderHandler{
		serv: *serv,
	}
}

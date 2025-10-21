package order

import (
	"order/internal/api"
	"order/internal/service"
	apii "order/pkg/api"
)

var _ api.OrderServer = (*Server)(nil)

type Server struct {
	apii.UnimplementedHandler
	serv service.OrderService
}

func NewOrderServer(serv service.OrderService) *Server {
	return &Server{
		serv: serv,
	}
}

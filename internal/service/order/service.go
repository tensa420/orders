package order

import (
	invClient "order/internal/client/grpc/inventory"
	paymClient "order/internal/client/grpc/payment"
	"order/internal/repository"
	"order/internal/service"
)

var _ service.OrderService = (*Service)(nil)

type Service struct {
	repository.OrderRepository
	invClient  invClient.Client
	paymClient paymClient.Client
}

func NewService(repo repository.OrderRepository, inventoryClient invClient.Client, paymClient paymClient.Client) *Service {
	return &Service{
		OrderRepository: repo,
		invClient:       inventoryClient,
		paymClient:      paymClient,
	}
}

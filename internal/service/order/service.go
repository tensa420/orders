package order

import (
	invClient "order/internal/client/grpc/inventory"
	paymClient "order/internal/client/grpc/payment"
	"order/internal/repository"
	"order/internal/service"
)

var _ service.OrderService = (*OrderService)(nil)

type OrderService struct {
	repo       repository.OrderRepository
	invClient  invClient.Client
	paymClient paymClient.Client
}

func NewOrderService(repo repository.OrderRepository, inventoryClient *invClient.Client, paymClient *paymClient.Client) *OrderService {
	return &OrderService{
		repo:       repo,
		invClient:  *inventoryClient,
		paymClient: *paymClient,
	}
}

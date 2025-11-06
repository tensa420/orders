package order

import (
	"order/internal/client/grpc"
	"order/internal/producer"
	"order/internal/repository"
	"order/internal/service"
)

var _ service.OrderService = (*OrderService)(nil)

type OrderService struct {
	repo       repository.OrderRepository
	invClient  grpc.InventoryClient
	paymClient grpc.PaymentClient
	prod       producer.OrderPaidProducer
}

func NewOrderService(repo repository.OrderRepository, prod producer.OrderPaidProducer, inventoryClient grpc.InventoryClient, paymClient grpc.PaymentClient) *OrderService {
	return &OrderService{
		repo:       repo,
		invClient:  inventoryClient,
		paymClient: paymClient,
		prod:       prod,
	}
}

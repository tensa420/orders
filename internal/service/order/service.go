package order

import (
	"order/internal/client/grpc"
	"order/internal/repository"
)

type service struct {
	repository.OrderRepository
	grpc.PaymentClient
	grpc.InventoryClient
}

func NewService(repo repository.OrderRepository, invClient grpc.InventoryClient, paymClient grpc.PaymentClient) *service {
	return &service{
		OrderRepository: repo,
		InventoryClient: invClient,
		PaymentClient:   paymClient,
	}
}

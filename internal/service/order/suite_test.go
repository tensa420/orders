package order

import (
	"context"
	"order/internal/service/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	ctx             context.Context
	orderService    *OrderService
	orderRepository *mocks.OrderRepository
	paymentClient   *mocks.PaymentClient
	inventoryClient *mocks.InventoryClient
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.orderRepository = new(mocks.OrderRepository)
	s.paymentClient = new(mocks.PaymentClient)
	s.inventoryClient = new(mocks.InventoryClient)
	s.orderService = NewOrderService(s.orderRepository, s.inventoryClient, s.paymentClient)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

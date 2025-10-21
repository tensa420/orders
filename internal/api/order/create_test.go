package order

import (
	"order/pkg/api"

	"github.com/google/uuid"
)

func (s *ServerSuite) TestCreateOrderServer_Success() {
	someOrderUUID := uuid.New()
	totalExample := 400.0
	resp := &api.CreateOrderResponse{
		OrderUUID:  someOrderUUID,
		TotalPrice: totalExample,
	}
	exampleUserUUID := uuid.New()
	req := &api.CreateOrderRequest{
		UserUUID:  exampleUserUUID,
		PartUuids: []string{"part1", "part2"},
	}

	s.service.On("CreateOrder", s.ctx, req.UserUUID.String(), req.PartUuids).Return(resp.OrderUUID.String(), resp.TotalPrice, nil).Once()

	result, err := s.server.HandleCreateOrder(s.ctx, req)

	s.NoError(err)
	s.Equal(resp, result)
	s.service.AssertExpectations(s.T())
}

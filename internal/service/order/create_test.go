package order

import (
	"order/internal/entity"
)

func (s *ServiceSuite) TestCreateOrder_Success() {
	ord := entity.Order{
		UserUUID:  "test-UUID",
		PartsUUID: []string{"part1", "part2"},
	}
	expectedOrderUUID := "testOrder-UUID"
	s.inventoryClient.On("ListParts", s.ctx, ord.PartsUUID).Return([]*entity.Part{
		{UUID: "part1", Price: 150},
		{UUID: "part2", Price: 250},
	}, nil).Once()

	s.orderRepository.On("CreateOrder", s.ctx, ord.UserUUID, ord.PartsUUID, 400.0).Return(expectedOrderUUID, nil).Once()

	result, total, err := s.orderService.CreateOrder(s.ctx, ord.UserUUID, ord.PartsUUID)

	s.NoError(err)
	s.Equal(expectedOrderUUID, result)
	s.Equal(400.0, total)
	s.inventoryClient.AssertExpectations(s.T())
	s.orderRepository.AssertExpectations(s.T())
}

package order

import (
	"order/internal/entity"
)

func (s *ServiceSuite) TestCancelOrder_Success() {
	orderUUID := "some-existing-uuid"
	order := &entity.Order{
		OrderUUID: orderUUID,
	}
	s.orderRepository.On("GetOrder", s.ctx, orderUUID).Return(order, nil).Once()
	s.orderRepository.On("CancelOrder", s.ctx, orderUUID).Return(nil).Once()

	err := s.orderService.CancelOrder(s.ctx, orderUUID)

	s.NoError(err)
	s.orderRepository.AssertExpectations(s.T())
}

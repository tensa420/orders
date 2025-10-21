package order

import "order/internal/entity"

func (s *ServiceSuite) TestGetOrder() {
	orderUUID := "test-uuid-123"
	expectedOrder := &entity.Order{
		OrderUUID: orderUUID,
		UserUUID:  "user-322",
		Status:    entity.StatusPaid,
	}
	s.orderRepository.On("GetOrder", s.ctx, orderUUID).Return(expectedOrder, nil).Once()
	order, err := s.orderService.GetOrder(s.ctx, orderUUID)

	s.NoError(err)
	s.Equal(expectedOrder, order)
}

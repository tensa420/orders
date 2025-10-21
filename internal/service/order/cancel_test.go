package order

import (
	"errors"
)

func (s *ServiceSuite) TestCancelOrder_NotFound() {
	orderUUID := "non-existing"
	expectedError := errors.New("order not found")

	s.orderRepository.On("CancelOrder", s.ctx, orderUUID).Return(expectedError).Once()

	err := s.orderService.CancelOrder(s.ctx, orderUUID)

	s.Error(err)
	s.orderRepository.AssertExpectations(s.T())
}

func (s *ServiceSuite) TestCancelOrder_Success() {
	orderUUID := "some-existing-uuid"

	s.orderRepository.On("CancelOrder", s.ctx, orderUUID).Return(nil).Once()

	err := s.orderService.CancelOrder(s.ctx, orderUUID)

	s.NoError(err)
	s.orderRepository.AssertExpectations(s.T())
}

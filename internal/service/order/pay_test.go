package order

import (
	"order/internal/client/converter"
	"order/internal/entity"
)

func (s *ServiceSuite) TestPayOrder() {
	orderUUID := "some-uuid"
	exampleOfPaymentMethod := "CARD"

	expectedOrder := entity.Order{
		OrderUUID: orderUUID,
		UserUUID:  "user-322",
	}
	s.orderRepository.On("GetOrder", s.ctx, orderUUID).Return(&expectedOrder, nil).Once()

	transUUID := "some-trans-uuid"
	s.paymentClient.On("PayOrder", s.ctx, expectedOrder.OrderUUID, expectedOrder.UserUUID, exampleOfPaymentMethod).Return(transUUID, nil).Once()

	paymentMethod := entity.PaymentMethodCard
	s.orderRepository.On("PayOrder", s.ctx, transUUID, expectedOrder.OrderUUID, paymentMethod).Return(nil).Once()

	result, err := s.orderService.PayOrder(s.ctx, orderUUID, converter.ConvertPaymentMethodToString(paymentMethod))

	s.NoError(err)
	s.Equal(result, transUUID)

	s.paymentClient.AssertExpectations(s.T())
	s.orderRepository.AssertExpectations(s.T())

}

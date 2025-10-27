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

	examplePaymentInfo := entity.PaymentInfo{
		TransactionUUID: "some-trans-uuid",
		OrderUUID:       orderUUID,
		PaymentMethod:   converter.ConvertPaymentMethod(exampleOfPaymentMethod),
	}
	s.paymentClient.On("PayOrder", s.ctx, expectedOrder.OrderUUID, expectedOrder.UserUUID, exampleOfPaymentMethod).Return(examplePaymentInfo.TransactionUUID, nil).Once()

	paymentMethod := entity.PaymentMethodCard
	s.orderRepository.On("PayOrder", s.ctx, examplePaymentInfo).Return(nil).Once()

	result, err := s.orderService.PayOrder(s.ctx, orderUUID, converter.ConvertPaymentMethodToString(paymentMethod))

	s.NoError(err)
	s.Equal(result, examplePaymentInfo.TransactionUUID)

	s.paymentClient.AssertExpectations(s.T())
	s.orderRepository.AssertExpectations(s.T())

}

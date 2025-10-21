package order

import (
	"order/internal/client/converter"
	"order/pkg/api"
)

func (s *ServerSuite) TestServerPay_Success() {
	req := &api.PayOrderRequest{
		PaymentMethod: "CARD",
	}
	params := api.HandlePayOrderParams{
		OrderUUID: converter.StringToUUID("00000000-0000-0000-0000-000000000001"),
	}
	someTransactionUUID := "00000000-0000-0000-0000-000000000002"
	s.service.On("PayOrder", s.ctx, params.OrderUUID.String(), req.PaymentMethod).Return(someTransactionUUID, nil).Once()

	result, err := s.server.HandlePayOrder(s.ctx, req, params)

	s.NoError(err)
	s.Equal(result, &api.PayOrderResponse{TransactionUUID: converter.StringToUUID(someTransactionUUID)})

	s.service.AssertExpectations(s.T())
}

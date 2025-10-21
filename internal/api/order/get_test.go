package order

import (
	"order/internal/client/converter"
	"order/internal/entity"
	"order/pkg/api"
)

func (s *ServerSuite) TestServerGet_Success() {
	params := api.HandleGetOrderParams{
		OrderUUID: converter.StringToUUID("00000000-0000-0000-0000-000000000001"),
	}

	someTransUUID := "00000000-0000-0000-0000-000000000006"
	pointToTransUUID := &someTransUUID

	somePaymentMethod := "00000000-0000-0000-0000-000000000007"
	pointToPaymMethod := &somePaymentMethod

	expectedOrder := &entity.Order{
		OrderUUID: "00000000-0000-0000-0000-000000000001",
		UserUUID:  "00000000-0000-0000-0000-000000000002",
		PartsUUID: []string{
			"00000000-0000-0000-0000-00000004",
			"00000000-0000-0000-0000-00000005",
		},
		TotalPrice:      400.0,
		TransactionUUID: pointToTransUUID,
		PaymentMethod:   pointToPaymMethod,
		Status:          "PAID",
	}
	s.service.On("GetOrder", s.ctx, params.OrderUUID.String()).Return(expectedOrder, nil).Once()

	result, err := s.server.HandleGetOrder(s.ctx, params)

	s.NoError(err)
	s.Equal(result, converter.ToAPI(expectedOrder))
}

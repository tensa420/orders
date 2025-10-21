package order

import (
	"order/internal/client/converter"
	"order/pkg/api"
)

func (s *ServerSuite) TestServerCancel_Success() {
	req := api.HandleCancelOrderParams{
		OrderUUID: converter.StringToUUID("00000000-0000-0000-0000-000000000001"),
	}
	s.service.On("CancelOrder", s.ctx, req.OrderUUID.String()).Return(nil).Once()

	response, err := s.server.HandleCancelOrder(s.ctx, req)

	s.NoError(err)
	s.NotNil(response)
	s.service.AssertExpectations(s.T())
}

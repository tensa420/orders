package order

import (
	"order/internal/entity"
	"reflect"

	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestCreateOrder_Success() {
	ord := entity.Order{
		UserUUID:   "test-UUID",
		PartsUUID:  []string{"part1", "part2"},
		TotalPrice: 400.0,
		Status:     entity.StatusPendingPayment,
	}

	s.inventoryClient.On("ListParts", s.ctx, ord.PartsUUID).Return([]*entity.Part{
		{UUID: "part1", Price: 150},
		{UUID: "part2", Price: 250},
	}, nil).Once()

	s.orderRepository.
		On("CreateOrder",
			mock.Anything,
			mock.MatchedBy(func(o entity.Order) bool {
				return o.UserUUID == ord.UserUUID &&
					reflect.DeepEqual(o.PartsUUID, ord.PartsUUID) &&
					o.TotalPrice == 400 &&
					o.Status == entity.StatusPendingPayment
			}),
		).
		Return(nil).
		Once()

	result, total, err := s.orderService.CreateOrder(s.ctx, ord.UserUUID, ord.PartsUUID)

	s.NoError(err)
	s.NotNil(result)
	s.Equal(400.0, total)
	s.inventoryClient.AssertExpectations(s.T())
	s.orderRepository.AssertExpectations(s.T())
}

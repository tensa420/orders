package order

import (
	"context"
	"order/internal/service/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetOrder(ctx context.Context, orderUUID string) (*model.GetOrderResponse, error) {
	req, err := s.OrderRepository.GetOrder(ctx, orderUUID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &model.GetOrderResponse{
		OrderUUID:       req.OrderUUID,
		UserUUID:        req.UserUUID,
		TransactionUUID: req.TransactionUUID,
		PaymentMethod:   req.PaymentMethod,
		PartUuids:       req.PartUuids,
		Status:          req.Status,
		TotalPrice:      req.TotalPrice,
	}, nil
}

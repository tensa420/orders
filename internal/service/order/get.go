package order

import (
	"context"
	repoModel "order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetOrder(ctx context.Context, orderUUID string) (*repoModel.GetOrderResponse, error) {
	req, err := s.repo.GetOrder(ctx, orderUUID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &repoModel.GetOrderResponse{
		OrderUUID:       req.OrderUUID,
		UserUUID:        req.UserUUID,
		TransactionUUID: req.TransactionUUID,
		PaymentMethod:   req.PaymentMethod,
		PartUuids:       req.PartUuids,
		Status:          req.Status,
		TotalPrice:      req.TotalPrice,
	}, nil
}

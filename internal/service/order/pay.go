package order

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) PayOrder(ctx context.Context, userUUID string, orderUUID string, paymentMethod string) (string, error) {
	req, err := s.PaymentClient.PayOrder(ctx, userUUID, orderUUID, paymentMethod)
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}
	err = s.OrderRepository.PayOrder(ctx, req, orderUUID)
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}
	return req, nil
}

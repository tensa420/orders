package order

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CancelOrder(ctx context.Context, orderUUID string) error {
	err := s.OrderRepository.CancelOrder(ctx, orderUUID)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return nil
}

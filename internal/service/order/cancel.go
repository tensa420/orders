package order

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) CancelOrder(ctx context.Context, orderUUID string) error {
	err := s.repo.CancelOrder(ctx, orderUUID)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return nil
}

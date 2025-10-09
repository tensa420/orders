package order

import (
	"context"
	"order/internal/client/converter"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) PayOrder(ctx context.Context, orderUUID string, paymentMethod string) (string, error) {
	transUUID := uuid.New()

	convertedPaymentMethod := converter.ConvertPaymentMethod(paymentMethod)
	err := s.OrderRepository.PayOrder(ctx, transUUID.String(), orderUUID, convertedPaymentMethod)
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}
	return transUUID.String(), nil
}

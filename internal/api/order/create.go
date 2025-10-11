package order

import (
	"context"
	ap "order/api"
	"order/internal/client/converter"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) HandleCreateOrder(ctx context.Context, req ap.CreateOrderRequest) (ap.HandleCreateOrderRes, error) {
	orderUUID, total, err := a.serv.CreateOrder(ctx, req.UserUUID.String(), req.PartUuids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &ap.CreateOrderResponse{
		OrderUUID:  converter.StringToUUID(orderUUID),
		TotalPrice: total,
	}, nil
}

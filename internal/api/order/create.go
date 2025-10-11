package order

import (
	"context"
	"order/internal/client/converter"
	"order/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Server) HandleCreateOrder(ctx context.Context, req *api.CreateOrderRequest) (api.HandleCreateOrderRes, error) {
	orderUUID, total, err := a.serv.CreateOrder(ctx, req.UserUUID.String(), req.PartUuids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &api.CreateOrderResponse{
		OrderUUID:  converter.StringToUUID(orderUUID),
		TotalPrice: total,
	}, nil
}

package order

import (
	"context"
	"order/internal/client/converter"
	"order/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Server) HandleGetOrder(ctx context.Context, params api.HandleGetOrderParams) (api.HandleGetOrderRes, error) {
	req, err := a.serv.GetOrder(ctx, params.OrderUUID.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &api.GetOrderResponse{
		OrderUUID:       converter.StringToUUID(req.OrderUUID),
		UserUUID:        converter.StringToUUID(req.UserUUID),
		PartUuids:       req.PartUuids,
		TotalPrice:      req.TotalPrice,
		TransactionUUID: converter.OptNilUUIDToUUID(req.TransactionUUID),
		PaymentMethod:   converter.OptNilStringToString(req.PaymentMethod),
		Status:          req.Status,
	}, nil
}

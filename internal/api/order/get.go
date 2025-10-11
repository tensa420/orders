package order

import (
	"context"
	ap "order/api"
	"order/internal/client/converter"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) HandleGetOrder(ctx context.Context, params ap.HandleGetOrderParams) (ap.HandleGetOrderRes, error) {
	req, err := a.serv.GetOrder(ctx, params.OrderUUID.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &ap.GetOrderResponse{
		OrderUUID:       converter.StringToUUID(req.OrderUUID),
		UserUUID:        converter.StringToUUID(req.UserUUID),
		PartUuids:       req.PartUuids,
		TotalPrice:      req.TotalPrice,
		TransactionUUID: converter.OptNilUUIDToUUID(req.TransactionUUID),
		PaymentMethod:   converter.OptNilStringToString(req.PaymentMethod),
		Status:          req.Status,
	}, nil
}

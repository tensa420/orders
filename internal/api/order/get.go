package order

import (
	"context"
	"order/internal/client/converter"
	"order/internal/entity"
	"order/pkg/api"
)

func (a *Server) HandleGetOrder(ctx context.Context, params api.HandleGetOrderParams) (api.HandleGetOrderRes, error) {
	req, err := a.serv.GetOrder(ctx, params.OrderUUID.String())
	if err != nil {
		return nil, entity.ErrOrderNotFound
	}
	return &api.GetOrderResponse{
		OrderUUID:       converter.StringToUUID(req.OrderUUID),
		UserUUID:        converter.StringToUUID(req.UserUUID),
		PartUuids:       req.PartsUUID,
		TotalPrice:      req.TotalPrice,
		TransactionUUID: converter.OptNilUUIDToUUID(req.TransactionUUID),
		PaymentMethod:   converter.OptNilStringToString(req.PaymentMethod),
		Status:          req.Status,
	}, nil
}

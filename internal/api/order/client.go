package order

import (
	apii "order/api"
	"order/internal/api"
	"order/internal/service"
	"order/pkg/inventory"
	"order/pkg/payment"
)

var _ api.Client = (*Api)(nil)

type Api struct {
	invent  inventory.InventoryServiceClient
	payment payment.PaymentClient
	serv    service.OrderService
}

func NewAPI(serv service.OrderService) *Api {

	return &Api{
		serv: serv,
	}
}

type OrderHandler struct {
	api Api
	apii.UnimplementedHandler
	paym payment.PaymentClient
	inv  inventory.InventoryServiceClient
}

func NewOrderHandler(api *Api) *OrderHandler {
	return &OrderHandler{
		api: *api,
	}
}

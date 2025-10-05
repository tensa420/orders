package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"sync"
)

import (
	ap "order/api"
	in "order/pkg/pb/inventory/inventory"
	pay "order/pkg/pb/payment/payment"
)

type Order struct {
	OrderUUID       string   `json:"order_uuid"`
	UserUUID        string   `json:"user_uuid"`
	PartsUUID       []string `json:"parts_uuid"`
	TotalPrice      float64  `json:"total_price"`
	TransactionUUID *string  `json:"transaction_uuid"`
	PaymentMethod   *string  `json:"payment_method"`
	Status          string   `json:"status"`
}

const (
	PaymentMethodUnknown = iota
	PaymentMethodCARD
	PaymentMethodSBP
	PaymentMethodCREDITCARD
	PaymentMethodINVESTORMONEY
)

var InventoryAddress = ":50052"
var PaymentAddress = ":50051"

type OrderStorage struct {
	mu     sync.RWMutex
	Orders map[string]*Order
}

func NewOrderStorage() *OrderStorage {
	return &OrderStorage{
		Orders: make(map[string]*Order),
	}
}

type OrderHandler struct {
	orders    *OrderStorage
	inventory in.InventoryServiceClient
	payment   pay.PaymentClient
}

func NewOrderHandler(order *OrderStorage, inv in.InventoryServiceClient, pay pay.PaymentClient) *OrderHandler {
	return &OrderHandler{orders: order, inventory: inv, payment: pay}
}

type Server struct {
	in.InventoryServiceClient
	pay.PaymentClient
}

func PaymToEnum(s string) (pay.PaymentMethod, error) {
	switch s {
	case "CARD":
		return pay.PaymentMethod_PAYMENT_METHOD_CARD, nil
	case "SBP":
		return pay.PaymentMethod_PAYMENT_METHOD_SBP, nil
	case "CREDITCARD":
		return pay.PaymentMethod_PAYMENT_METHOD_CARD, nil
	case "INVESTORMONEY":
		return pay.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY, nil
	default:
		return pay.PaymentMethod_PAYMENT_METHOD_UNKNOWN, nil
	}
}
func (s *OrderHandler) CreateOrders(ctx context.Context, reqBody ap.CreateOrderRequest, server Server) (*ap.CreateOrderResponse, error) {

	grpcToIn := &in.ListPartsRequest{
		Filter: &in.PartsFilter{
			Uuids: reqBody.PartUuids,
		},
	}

	grpcFromIn, err := server.ListParts(ctx, grpcToIn)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var TotalPrice float64

	for _, part := range grpcFromIn.Parts {
		reqToGetPart := &in.GetPartRequest{Uuid: part.UUID}
		_, err1 := server.GetPart(ctx, reqToGetPart)
		if err1 != nil {
			return nil, status.Error(codes.Internal, err1.Error())
		}
		TotalPrice += part.Price
	}
	UUID := uuid.New()
	order := &Order{
		OrderUUID:  UUID.String(),
		UserUUID:   reqBody.UserUUID.String(),
		PartsUUID:  reqBody.PartUuids,
		TotalPrice: TotalPrice,
		Status:     "PENDING_PAYMENT",
	}

	s.orders.Orders[order.OrderUUID] = order

	return &ap.CreateOrderResponse{
		OrderUUID:  UUID,
		TotalPrice: TotalPrice,
	}, nil
}

func (s *OrderHandler) PayOrder(ctx context.Context, res ap.PayOrderRequest, params ap.HandlePayOrderParams, server Server) (*ap.PayOrderResponse, error) {

	ord, ok := s.orders.Orders[params.OrderUUID.String()]
	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	pm, err := PaymToEnum(res.PaymentMethod)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	grpcToPay := &pay.PayOrderRequest{
		Order: &pay.OrderRequest{
			OrderUuid:     ord.OrderUUID,
			UserUuid:      ord.UserUUID,
			PaymentMethod: pm,
		},
	}
	grpcFromPay, err := server.PayOrder(ctx, grpcToPay)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.orders.Orders[params.OrderUUID.String()].Status = "PAID"
	s.orders.Orders[params.OrderUUID.String()].TransactionUUID = &grpcFromPay.TransactionUuid

	transUUID, err := uuid.Parse(grpcFromPay.TransactionUuid)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	return &ap.PayOrderResponse{
		TransactionUUID: transUUID,
	}, nil
}

func (s *OrderHandler) GetOrder(params ap.HandleGetOrderParams, server Server) (*ap.GetOrderResponse, error) {
	ord, ok := s.orders.Orders[params.OrderUUID.String()]
	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	parsedUserUUID, err := uuid.Parse(ord.UserUUID)
	transUUID, err := uuid.Parse(*ord.TransactionUUID)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	finalTransUUID := ap.OptNilUUID{Value: transUUID, Set: true, Null: false}
	finalPaymentMethod := ap.OptNilString{Value: *ord.PaymentMethod, Set: true, Null: false}

	return &ap.GetOrderResponse{
		OrderUUID:       params.OrderUUID,
		UserUUID:        parsedUserUUID,
		PartUuids:       ord.PartsUUID,
		TotalPrice:      ord.TotalPrice,
		TransactionUUID: finalTransUUID,
		PaymentMethod:   finalPaymentMethod,
		Status:          ord.Status,
	}, nil
}

func (s *OrderHandler) CancelOrder(params ap.HandleCancelOrderParams) ap.Error {

	ord, ok := s.orders.Orders[params.OrderUUID.String()]
	if !ok {
		return ap.Error{Code: "404", Message: "order not found"}
	}
	if ord.Status == "PAID" {
		return ap.Error{
			Code:    "409",
			Message: "order already paid",
		}
	}
	if ord.Status == "PENDING_PAYMENT" {
		ord.Status = "CANCELLED"
		return ap.Error{
			Code:    "204",
			Message: "No content",
		}
	}
	return ap.Error{}
}
func main() {
	orders := NewOrderStorage()

}

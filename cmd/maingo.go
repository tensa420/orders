package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	ap "order/api"
	in "order/pkg/pb/inventory/inventory"
	pay "order/pkg/pb/payment/payment"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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
func (s *OrderHandler) HandleCreateOrder(ctx context.Context, req *ap.CreateOrderRequest) (ap.HandleCreateOrderRes, error) {

	grpcToIn := &in.ListPartsRequest{
		Filter: &in.PartsFilter{
			Uuids: req.PartUuids,
		},
	}

	grpcFromIn, err := s.inventory.ListParts(ctx, grpcToIn)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var total float64
	for _, part := range grpcFromIn.Parts {
		reqToGetPart := &in.GetPartRequest{Uuid: part.UUID}
		_, err1 := s.inventory.GetPart(ctx, reqToGetPart)
		if err1 != nil {
			return nil, status.Error(codes.Internal, err1.Error())
		}
		total += part.Price
	}

	UUID := uuid.New()

	order := &Order{
		OrderUUID:  UUID.String(),
		UserUUID:   req.UserUUID.String(),
		PartsUUID:  req.PartUuids,
		TotalPrice: total,
		Status:     "PENDING_PAYMENT",
	}

	s.orders.mu.Lock()
	s.orders.Orders[order.OrderUUID] = order
	s.orders.mu.Unlock()

	return &ap.CreateOrderResponse{
		OrderUUID:  UUID,
		TotalPrice: total,
	}, nil
}

func (s *OrderHandler) HandlePayOrder(ctx context.Context, req *ap.PayOrderRequest, params ap.HandlePayOrderParams) (ap.HandlePayOrderRes, error) {
	s.orders.mu.RLock()
	ord, ok := s.orders.Orders[params.OrderUUID.String()]
	s.orders.mu.RUnlock()

	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	pm, err := PaymToEnum(req.PaymentMethod)
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
	grpcFromPay, err := s.payment.PayOrder(ctx, grpcToPay)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.orders.mu.Lock()
	s.orders.Orders[params.OrderUUID.String()].Status = "PAID"
	s.orders.Orders[params.OrderUUID.String()].TransactionUUID = &grpcFromPay.TransactionUuid
	s.orders.mu.Unlock()

	transUUID, err := uuid.Parse(grpcFromPay.TransactionUuid)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	return &ap.PayOrderResponse{
		TransactionUUID: transUUID,
	}, nil
}

func (s *OrderHandler) HandleGetOrder(ctx context.Context, params ap.HandleGetOrderParams) (ap.HandleGetOrderRes, error) {
	s.orders.mu.RLock()
	ord, ok := s.orders.Orders[params.OrderUUID.String()]
	s.orders.mu.RUnlock()
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

func (s *OrderHandler) HandleCancelOrder(ctx context.Context, params ap.HandleCancelOrderParams) (ap.HandleCancelOrderRes, error) {

	ord, ok := s.orders.Orders[params.OrderUUID.String()]
	if !ok {
		return &ap.HandleCancelOrderNotFound{Code: "NotFound", Message: "Order not found"}, nil
	}

	if ord.Status == "PAID" {
		return &ap.HandleCancelOrderConflict{Code: "CancelOrderConflict", Message: "Order is already paid"}, nil
	}

	ord.Status = "CANCELLED"
	return &ap.HandleCancelOrderNoContent{}, nil
}
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	connInventory, err := grpc.DialContext(
		ctx,
		InventoryAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to inventory: %v", err)
	}

	inventoryClient := in.NewInventoryServiceClient(connInventory)
	defer connInventory.Close()

	connPayment, err := grpc.DialContext(
		ctx,
		PaymentAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to payment: %v", err)
	}

	paymentClient := pay.NewPaymentClient(connPayment)
	defer connPayment.Close()

	orders := NewOrderStorage()
	hand := NewOrderHandler(orders, inventoryClient, paymentClient)
	orderHandler, err := ap.NewServer(hand, nil)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: orderHandler,
	}
	go func() {
		serv, err := net.Listen("tcp", srv.Addr)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		err = srv.Serve(serv)
		if err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server shutdown successfully")
}

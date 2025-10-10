package main

import (
	"context"
	"log"
	"net"
	"net/http"
	ap "order/internal/api/order"
	clientInv "order/internal/client/grpc/inventory"
	clientPaym "order/internal/client/grpc/payment"
	repo "order/internal/repository/order"
	service "order/internal/service/order"
	"order/pkg/inventory"
	payment "order/pkg/payment"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var InventoryAddress = ":50052"
var PaymentAddress = ":50051"

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

	connPayment, err := grpc.DialContext(
		ctx,
		PaymentAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to payment: %v", err)
	}
	generatedInvClient := inventory.NewInventoryServiceClient(connInventory)
	inventoryClient := clientInv.NewClient(generatedInvClient)
	generatedPaymClient := payment.NewPaymentClient(connPayment)
	paymentClient := clientPaym.New(generatedPaymClient)

	defer connInventory.Close()
	defer connPayment.Close()

	orders := repo.NewRepository()
	serv := service.NewService(orders, inventoryClient, paymentClient)
	api := ap.NewAPI(serv, inventoryClient, paymentClient)
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

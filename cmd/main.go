package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"order/pkg/inventory"
	"order/pkg/payment"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	ap "order/internal/api/order"
	clientInv "order/internal/client/grpc/inventory"
	clientPaym "order/internal/client/grpc/payment"
	repo "order/internal/repository/repository"
	service "order/internal/service/order"
	apii "order/pkg/api"
)

var (
	InventoryAddress = "localhost:50062"
	PaymentAddress   = ":50051"
)

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
		log.Printf("Failed to connect to inventory: %v", err)
		return
	}

	connPayment, err := grpc.DialContext(
		ctx,
		PaymentAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("Failed to connect to payment: %v", err)
		return
	}

	generatedInvClient := inventory.NewInventoryServiceClient(connInventory)
	inventoryClient := clientInv.NewClient(generatedInvClient)
	generatedPaymClient := payment.NewPaymentClient(connPayment)
	paymentClient := clientPaym.New(generatedPaymClient)

	defer connInventory.Close()
	defer connPayment.Close()

	con, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = con.Ping(ctx)
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	orderRepo := repo.NewOrderRepository(con)
	orderSerivce := service.NewOrderService(orderRepo, inventoryClient, paymentClient)
	orderServer := ap.NewOrderServer(orderSerivce)

	orderHandler, err := apii.NewServer(orderServer)
	if err != nil {
		log.Printf("Failed to create server: %v", err)
		return
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: orderHandler,
	}

	go func() {
		serv, err := net.Listen("tcp", srv.Addr)
		if err != nil {
			log.Printf("Failed to listen: %v", err)
		}
		err = srv.Serve(serv)
		if err != nil {
			log.Printf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Server shutdown successfully")
}

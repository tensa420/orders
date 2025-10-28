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

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
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
	inventoryAddr = "inventoryService:50062"
	paymentAddr   = "paymentService:50051"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	connInventory, err := grpc.DialContext(
		ctx,
		inventoryAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("Failed to connect to inventory: %v", err)
		return
	}

	connPayment, err := grpc.DialContext(
		ctx,
		paymentAddr,
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

	con, err := pgx.Connect(ctx, os.Getenv("DB_URI"))
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return
	}

	err = con.Ping(ctx)
	if err != nil {
		log.Printf("Failed to ping DB: %v", err)
		return
	}

	migrations := repo.NewMigrator(stdlib.OpenDB(*con.Config().Copy()), "migrations")
	err = migrations.Up()
	if err != nil {
		log.Printf("Failed to apply migrations: %v", err)
		return
	}

	err = con.Close(ctx)
	if err != nil {
		log.Printf("Failed to close connection with db: %v", err)
	}

	pool, err := pgxpool.New(ctx, os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	orderRepo := repo.NewOrderRepository(pool)
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

package app

import (
	"context"
	"log"
	"net"
	api2 "order/internal/api"
	api "order/internal/api/order"
	inventory2 "order/internal/client/grpc/inventory"
	payment2 "order/internal/client/grpc/payment"
	"order/internal/repository"
	repository2 "order/internal/repository/repository"
	"order/internal/service"
	"order/internal/service/order"
	"order/pkg/inventory"
	"order/pkg/payment"
	"order/platform/pkg/closer"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type diContainer struct {
	orderServ *api2.OrderServer

	orderUseCase service.OrderService

	orderRepository repository.OrderRepository

	pool *pgxpool.Pool

	inventoryServiceClient *inventory2.Client

	paymentServiceClient *payment2.Client
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderApi(ctx context.Context) api2.OrderServer {
	if d.orderServ == nil {
		return api.NewOrderServer(d.OrderService(ctx))
	}
	return *d.orderServ
}

func (d *diContainer) OrderService(ctx context.Context) service.OrderService {
	if d.orderUseCase == nil {
		return order.NewOrderService(d.OrderRepository(ctx), d.InventoryClient(ctx), d.PaymentClient(ctx))
	}
	return d.orderUseCase
}

func (d *diContainer) OrderRepository(ctx context.Context) repository.OrderRepository {
	if d.orderRepository == nil {
		return repository2.NewOrderRepository(d.Pool(ctx))
	}
	return d.orderRepository
}

func (d *diContainer) Pool(ctx context.Context) *pgxpool.Pool {
	if d.pool == nil {
		err := godotenv.Load("./deploy/env/.env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		pool, err := pgxpool.New(ctx, os.Getenv("DB_URI"))
		if err != nil {
			log.Fatalf("Failed to close connection with db: %v", err)
		}
		return pool
	}

	return d.pool
}

func (d *diContainer) InventoryClient(ctx context.Context) *inventory2.Client {
	if d.inventoryServiceClient != nil {
		return d.inventoryServiceClient
	}
	err := godotenv.Load("./deploy/env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	invHost := os.Getenv("INVENTORY_CLIENT_HOST")
	invPort := os.Getenv("INVENTORY_CLIENT_PORT")

	inventoryConn, err := CreateGRPCConnection(net.JoinHostPort(invHost, invPort))
	if err != nil {
		log.Println(net.JoinHostPort(invHost, invPort))
		log.Fatalf("Error connecting to inventory service: %v", err)
	}

	closer.AddNamed("inventory client connection", func(context.Context) error {
		return inventoryConn.Close()
	})

	generatedClient := inventory.NewInventoryServiceClient(inventoryConn)
	d.inventoryServiceClient = inventory2.NewClient(generatedClient)

	return d.inventoryServiceClient
}

func (d *diContainer) PaymentClient(ctx context.Context) *payment2.Client {
	if d.paymentServiceClient == nil {
		err := godotenv.Load("./deploy/env/.env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		paymHost := os.Getenv("PAYMENT_CLIENT_HOST")
		paymPort := os.Getenv("PAYMENT_CLIENT_PORT")
		paymentConn, err := CreateGRPCConnection(net.JoinHostPort(paymHost, paymPort))

		if err != nil {
			log.Println(net.JoinHostPort(paymHost, paymPort))
			log.Fatalf("Error connecting to payment service: %v", err)
		}

		closer.AddNamed("payment client connection", func(context.Context) error {
			return paymentConn.Close()
		})
		generatedClient := payment.NewPaymentClient(paymentConn)
		d.paymentServiceClient = payment2.New(generatedClient)

		return d.paymentServiceClient
	}
	return d.paymentServiceClient
}

func CreateGRPCConnection(addr string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

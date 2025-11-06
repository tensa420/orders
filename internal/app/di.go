package app

import (
	"context"
	"log"
	"net"
	api2 "order/internal/api"
	api "order/internal/api/order"
	inventory2 "order/internal/client/grpc/inventory"
	payment2 "order/internal/client/grpc/payment"
	"order/internal/consumer"
	consumer2 "order/internal/consumer/consumer"
	"order/internal/producer"
	"order/internal/producer/order_paid"
	"order/internal/repository"
	repository2 "order/internal/repository/repository"
	"order/internal/service"
	"order/internal/service/order"
	"order/pkg/inventory"
	"order/pkg/payment"
	"order/platform/pkg/closer"
	"os"
	"time"

	"github.com/IBM/sarama"
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

	consumer sarama.ConsumerGroup
	producer sarama.SyncProducer

	orderPaidProd     producer.OrderPaidProducer
	shipAssembled1    consumer.ShipAssembledConsumer
	shipAssembledCons consumer.ShipAssembledConsumerService
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
		return order.NewOrderService(d.OrderRepository(ctx), d.NewOrderPaidProducer(ctx), d.InventoryClient(ctx), d.PaymentClient(ctx))
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

func (d *diContainer) Producer(ctx context.Context) sarama.SyncProducer {
	cfg := ProducerConfig()
	prod, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_BROKER")}, cfg)
	if err != nil {
		log.Printf("Error creating sync producer: %v", err)
		return nil
	}
	return prod
}

func (d *diContainer) ConsumerGroup(ctx context.Context) sarama.ConsumerGroup {
	cfg := ConsumerConfig()
	cons, err := sarama.NewConsumerGroup([]string{os.Getenv("KAFKA_BROKER")}, "EASY-KAFKA-GROUP", cfg)
	if err != nil {
		log.Printf("Error creating consumer group: %v", err)
		return nil
	}
	return cons
}

func (d *diContainer) NewShipAssembledConsumer(ctx context.Context) consumer.ShipAssembledConsumerService {
	d.shipAssembledCons = consumer2.NewShipAssembledConsumer(d.ConsumerGroup(ctx), []string{os.Getenv("KAFKA_CONSUMER_TOPIC")})
	return d.shipAssembledCons
}

func (d *diContainer) NewOrderPaidProducer(ctx context.Context) producer.OrderPaidProducer {
	d.orderPaidProd = order_paid.NewOrderPaidProducer(d.Producer(ctx), os.Getenv("KAFKA_PRODUSER_TOPIC"))
	return d.orderPaidProd
}

func ConsumerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}

func ProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Producer.Return.Successes = true

	return config
}

func (d *diContainer) NewShipAssembledCons(ctx context.Context) consumer.ShipAssembledConsumer {
	d.shipAssembled1 = consumer2.NewShipAssembledConsumer(d.ConsumerGroup(ctx), []string{os.Getenv("KAFKA_CONSUMER_TOPIC")})
	return d.shipAssembled1
}

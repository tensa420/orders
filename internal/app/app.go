package app

import (
	"context"
	"log"
	"net"
	"net/http"
	repo "order/internal/repository/repository"
	"os"

	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"order/internal/config"
	"order/pkg/api"
	"order/platform/pkg/closer"
	"order/platform/pkg/logger"
)

type App struct {
	diContainer *diContainer
	httpServer  *api.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}
func (a *App) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		logger.Info(ctx, "http server starting...")
		return a.RunHTTPServer(ctx)
	})

	g.Go(func() error {
		logger.Info(ctx, "consumer starting...")
		return a.RunConsumer(ctx)
	})

	// Здесь Run будет блокировать main, пока:
	// - одна из горутин не вернёт ошибку
	// - или ctx не отменят по сигналу
	return g.Wait()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.InitLogger,
		a.InitCloser,
		a.InitMigrator,
		a.InitHTTPServer,
	}
	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initDI(ctx context.Context) error {
	a.diContainer = NewDIContainer()
	return nil
}
func (a *App) InitLogger(ctx context.Context) error {
	return logger.Init(config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson())
}
func (a *App) InitCloser(ctx context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}
func (a *App) InitHTTPServer(ctx context.Context) error {
	var err error
	a.httpServer, err = api.NewServer(a.diContainer.OrderApi(ctx))
	if err != nil {
		return err
	}
	return nil
}

func (a *App) RunHTTPServer(ctx context.Context) error {
	orderServer := a.diContainer.OrderApi(ctx)

	orderHandler, err := api.NewServer(orderServer)
	if err != nil {
		logger.Error(ctx, "Failed to create order handler:", zap.Error(err))
	}
	srv := &http.Server{
		Addr:    config.AppConfig().OrderHTTP.Address(),
		Handler: orderHandler,
	}
	lis, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		logger.Error(ctx, "Failed to listen server: ", zap.Error(err))
	}
	logger.Info(ctx, "server started")
	err = srv.Serve(lis)
	if err != nil {
		logger.Error(ctx, "Failed to serve server: ", zap.Error(err))
	}
	return nil
}

func (a *App) InitMigrator(ctx context.Context) error {
	err := godotenv.Load("./deploy/env/.env")
	if err != nil {
		logger.Error(ctx, "Failed to load .env file", zap.Error(err))
	}

	con, err := pgx.Connect(ctx, os.Getenv("DB_URI"))
	if err != nil {
		logger.Error(ctx, "Failed to connect to database:", zap.Error(err))
		return err
	}

	closer.AddNamed("connect to postgres", con.Close)

	err = con.Ping(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to ping db:", zap.Error(err))
		return err
	}

	migrations := repo.NewMigrator(stdlib.OpenDB(*con.Config().Copy()), "./migrations")
	err = migrations.Up()
	if err != nil {
		logger.Error(ctx, "Failed to run migrations:", zap.Error(err))
		return err
	}

	return nil
}

func (a *App) RunConsumer(ctx context.Context) error {
	log.Printf("Consumer started!!!")

	service := a.diContainer.NewShipAssembledConsumer(ctx)
	if service == nil {
		return errors.New("shipAssembledCons is nil (check Kafka consumer init)")
	}

	handler := a.diContainer.NewShipAssembledCons(ctx)
	if handler == nil {
		return errors.New("shipAssembled1 is nil (consumer handler not inited)")
	}

	// 3. Запускаем
	err := service.RunConsumer(
		ctx,
		[]string{os.Getenv("KAFKA_CONSUMER_TOPIC")},
		handler,
	)
	if err != nil {
		log.Printf("Failed to run consumer: %v", err)
		return err
	}
	return nil
}

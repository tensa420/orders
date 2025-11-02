package main

import (
	"context"
	"log"
	"order/internal/app"
	"order/internal/config"
	"order/platform/pkg/closer"
	"order/platform/pkg/logger"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

const envPath = "./deploy/env/.env"

func main() {
	err := config.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer cancel()
	defer gracefulShutdown()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	app, err := app.NewApp(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to create app", zap.Error(err))
		return
	}

	err = app.Run(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to run app", zap.Error(err))
		return
	}
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := closer.CloseAll(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to close all graceful shutdown", zap.Error(err))
		return
	}
}

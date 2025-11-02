package config

import (
	"order/internal/config/env"
	"os"

	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Logger    LoggerConfig
	OrderHTTP OrderHTTPConfig
	Postgres  PostgresConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	orderHTTPCfg, err := env.NewOrderHTTPConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:    loggerCfg,
		OrderHTTP: orderHTTPCfg,
		Postgres:  postgresCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

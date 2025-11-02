package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type OrderHTTPEnvConfig struct {
	Host string `env:"HTTP_HOST,required"`
	Port string `env:"HTTP_PORT,required"`
}

type OrderHTTPConfig struct {
	config OrderHTTPEnvConfig
}

func NewOrderHTTPConfig() (*OrderHTTPConfig, error) {
	var config OrderHTTPEnvConfig
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return &OrderHTTPConfig{config: config}, nil
}

func (cfg *OrderHTTPConfig) Address() string {
	return net.JoinHostPort(cfg.config.Host, cfg.config.Port)
}

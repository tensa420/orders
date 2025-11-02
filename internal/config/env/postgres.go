package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type postgresEnvConfig struct {
	Host     string `env:"POSTGRES_HOST,required"`
	Port     string `env:"POSTGRES_PORT,required"`
	User     string `env:"POSTGRES_INITDB_ROOT_USERNAME,required"`
	Password string `env:"POSTGRES_INITDB_ROOT_PASSWORD,required"`
	Database string `env:"POSTGRES_DATABASE,required"`
}
type postgresConfig struct {
	cfg postgresEnvConfig
}

func NewPostgresConfig() (*postgresConfig, error) {
	var cfg postgresEnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &postgresConfig{cfg}, nil
}

func (cfg *postgresConfig) URI() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.cfg.User,
		cfg.cfg.Password,
		cfg.cfg.Host,
		cfg.cfg.Port,
		cfg.cfg.Database,
	)
}

func (cfg *postgresConfig) DatabaseName() string {
	return cfg.cfg.Database
}

package config

import (
	"crudl/internal/adapters/cache"
	kafkaconsumer "crudl/internal/adapters/kafka_consumer"
	"crudl/internal/adapters/postgres"
	"crudl/internal/adapters/postgres/migrations"
	"crudl/pkg/http_server"
	"crudl/pkg/logger"
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Migrate  migrations.Config
	Postgres postgres.Config
	Cache    cache.Config
	Kafka    kafkaconsumer.Config
	Http     http_server.Config
	Logger   logger.Config
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return &Config{}, errors.New("Error loading .env file")
	}

	err := envconfig.Process("", &cfg)
	if err != nil {
		return &Config{}, errors.New("Error value: " + err.Error())
	}

	return &cfg, nil
}

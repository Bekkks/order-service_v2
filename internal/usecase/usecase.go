package usecase

import (
	"context"
	"crudl/internal/adapters/cache"
	kafkaconsumer "crudl/internal/adapters/kafka_consumer"
	"crudl/internal/adapters/postgres"
	"crudl/internal/domain"

	"github.com/segmentio/kafka-go"
)

type Postgres interface {
	CreateOrder(ctx context.Context, sub domain.Order) error
	GetOrder(ctx context.Context, user_id string) (domain.Order, error)
}

type Cache interface {
	Get(orderUID string) (domain.Order, bool)
	Add(order domain.Order)
}

type Kafka interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
}

type Profile struct {
	Postgres Postgres
	Cache    Cache
	Kafka    Kafka
}

func NewProfile(postgres *postgres.Pool, cache *cache.Cache, kafka *kafkaconsumer.Consumer) *Profile {
	return &Profile{
		Postgres: postgres,
		Cache:    cache,
		Kafka: kafka,
	}
}

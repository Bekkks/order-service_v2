package usecase

import (
	"context"
	"crudl/internal/domain"
	"crudl/pkg/logger"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func (p *Profile) CreateOrder(ctx context.Context) error {
	msgChan := p.ConsumerRead(ctx)

	logger.Info("kafka consumer started")

	for m := range msgChan {
		var data domain.Order
		if err := json.Unmarshal(m.Value, &data); err != nil {
			logger.Error("failed to unmarshal kafka message", err)
			continue 
		}
		
		logger.Info("received message from kafka", "message", data.OrderUID)

		if err := data.Validate(); err != nil{
			return fmt.Errorf("validation error: %w", err)
		}

		if err := p.Postgres.CreateOrder(ctx, data); err != nil {
			logger.Error("failed to create subscription", err)
		}else{
			p.Cache.Add(data)
		}
	}

	return nil
}

func (p *Profile) ConsumerRead(ctx context.Context) <-chan kafka.Message {
	msg := make(chan kafka.Message, 100)

	go func() {
		defer close(msg)

		for {
			m, err := p.Kafka.ReadMessage(ctx)
			if err != nil {
				logger.Error("failed to read message from kafka", err)
				return
			}
			msg <- m
		}
	}()

	return msg
}


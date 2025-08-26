package cache

import (
	"context"
	"fmt"

	"crudl/internal/adapters/postgres"
	"crudl/internal/domain"

	lru "github.com/hashicorp/golang-lru/v2"
)

type Config struct {
	Size int `envconfig:"CACHE_SIZE" default:"10"`
}

type Cache struct {
	lru *lru.Cache[string, domain.Order]
}

func New(ctx context.Context, cfg Config, postgres *postgres.Pool) (*Cache, error) {
	lruCache, err := lru.New[string, domain.Order](cfg.Size)
	if err != nil {
		return nil, fmt.Errorf("Error creating LRU cache: %w", err)
	}

	rows, err := postgres.DB.QueryContext(ctx, "SELECT order_uid FROM orders LIMIT $1", cfg.Size)
	if err != nil {
		return nil, fmt.Errorf("Error fetching initial data for cache: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var orderUid string
		if err := rows.Scan(&orderUid); err != nil {
			return nil, fmt.Errorf("Error scanning row for cache: %w", err)
		}

		order, err := postgres.GetOrder(ctx, orderUid)
		if err != nil {
			return nil, fmt.Errorf("Error fetching order from DB: %w", err)
		}
		lruCache.Add(orderUid, order)
	}

	return &Cache{lru: lruCache}, nil
}
package postgres

import (
	"context"
	"crudl/internal/domain"
	"fmt"
)

func (p *Pool) GetOrder(ctx context.Context, order_id string) (domain.Order, error) {
	var sub domain.Order

	query := `SELECT *
			FROM orders
			WHERE orders.order_uid = $1`

	if err := p.DB.GetContext(ctx, &sub, query, order_id); err != nil {
		return domain.Order{}, fmt.Errorf("Error in get orders: %w", err)
	}

	query = `SELECT *
		FROM deliverys
		WHERE deliverys.order_uid = $1`

	if err := p.DB.GetContext(ctx, &sub.Delivery, query, order_id); err != nil {
		return domain.Order{}, fmt.Errorf("Error in get orders: %w", err)
	}

	query = `SELECT *
		FROM payments
		WHERE payments.order_uid = $1`

	if err := p.DB.GetContext(ctx, &sub.Payment, query, order_id); err != nil {
		return domain.Order{}, fmt.Errorf("Error in get orders: %w", err)
	}

	query = `SELECT *
		FROM items
		WHERE items.order_uid = $1`

	if err := p.DB.SelectContext(ctx, &sub.Items, query, order_id); err != nil {
		return domain.Order{}, fmt.Errorf("Error in get items: %w", err)
	}

	return sub, nil
}
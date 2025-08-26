package usecase

import (
	"context"
	"crudl/internal/domain"
)

func (p *Profile) GetOrder(ctx context.Context, order_id string) (domain.Order, error){
	if order, ok := p.Cache.Get(order_id); ok {
		return order, nil
	}

	subs, err := p.Postgres.GetOrder(ctx, order_id)
	if err != nil{
		return domain.Order{}, err
	}

	return subs, nil
}

package cache

import "crudl/internal/domain"

func (c *Cache) Add(order domain.Order) {
	c.lru.Add(order.OrderUID, order)
}
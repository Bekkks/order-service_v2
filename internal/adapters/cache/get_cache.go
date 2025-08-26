package cache

import "crudl/internal/domain"

func (c *Cache) Get(orderUID string) (domain.Order, bool) {
	return c.lru.Get(orderUID)
}
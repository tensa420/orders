package order

import (
	"order/internal/repository"
	repomModel "order/internal/repository/model"
	"sync"
)

var _ repository.OrderRepository = (*OrderRepository)(nil)

type OrderRepository struct {
	mu     sync.RWMutex
	orders map[string]repomModel.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]repomModel.Order),
	}
}

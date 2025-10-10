package order

import (
	"order/internal/repository"
	repomModel "order/internal/repository/model"
	v1 "order/pkg/inventory"
	"sync"
)

var _ repository.OrderRepositoryRepo = (*Repository)(nil)

type Repository struct {
	mu     sync.RWMutex
	inv    v1.InventoryServiceClient
	orders map[string]repomModel.Order
}

func NewRepository() *Repository {
	return &Repository{
		orders: make(map[string]repomModel.Order),
	}
}

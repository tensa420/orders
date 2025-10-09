package order

import (
	model "order/internal/repository/model"
	v1 "order/pkg/inventory/inventory"
	"sync"
)

type repository struct {
	mu     sync.RWMutex
	inv    v1.InventoryServiceClient
	orders map[string]model.Order
}

func NewRepository() *repository {
	return &repository{
		orders: make(map[string]model.Order),
	}
}

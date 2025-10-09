package inventory

import (
	inv "order/pkg/inventory/inventory"
)

type client struct {
	generatedClient inv.InventoryServiceClient
}

func NewClient(generatedClient inv.InventoryServiceClient) *client {
	return &client{generatedClient: generatedClient}
}

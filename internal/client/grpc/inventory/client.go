package inventory

import (
	inv "order/pkg/inventory"
)

type Client struct {
	generatedClient inv.InventoryServiceClient
}

func NewClient(generatedClient inv.InventoryServiceClient) *Client {
	return &Client{generatedClient: generatedClient}
}

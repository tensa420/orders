package inventory

import (
	"context"
	"order/internal/client/converter"
	"order/internal/entity"
	v1 "order/pkg/inventory"
)

func (c *Client) ListParts(ctx context.Context, partsUUID []string) ([]*entity.Part, error) {
	parts, err := c.generatedClient.ListParts(ctx, &v1.ListPartsRequest{
		Filter: &v1.PartsFilter{
			Uuids: partsUUID,
		},
	})
	if err != nil {
		return nil, err
	}
	return converter.PartsListToEntity(parts.Parts), nil
}

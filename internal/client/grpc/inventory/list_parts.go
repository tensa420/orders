package inventory

import (
	"context"
	"order/internal/client/converter"
	repoModel "order/internal/repository/model"
	v1 "order/pkg/inventory"
)

func (c *Client) ListParts(ctx context.Context, partsUUID []string) ([]*repoModel.Part, error) {
	parts, err := c.generatedClient.ListParts(ctx, &v1.ListPartsRequest{
		Filter: &v1.PartsFilter{
			Uuids: partsUUID,
		},
	})
	if err != nil {
		return nil, err
	}
	return converter.PartsListToRepoModel(parts.Parts), nil
}

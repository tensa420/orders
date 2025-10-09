package inventory

import (
	"context"
	"order/internal/client/converter"
	"order/internal/repository/model"
	v1 "order/pkg/inventory/inventory"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) ([]*model.Part, error) {
	parts, err := c.generatedClient.ListParts(ctx, &v1.ListPartsRequest{
		Filter: converter.PartFilterToProto(filter),
	})
	if err != nil {
		return nil, err
	}
	return converter.PartsListToModel(parts.Parts), nil
}

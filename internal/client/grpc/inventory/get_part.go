package inventory

import (
	"context"
	"order/internal/client/converter"
	"order/internal/repository/model"
	v1 "order/pkg/inventory/inventory"
)

func (c *client) GetPart(ctx context.Context, partUUID string) (*model.Part, error) {
	req, err := c.generatedClient.GetPart(ctx, &v1.GetPartRequest{
		Uuid: partUUID,
	})
	if err != nil {
		return nil, err
	}
	return converter.PartProtoToModel(req.Part), nil
}

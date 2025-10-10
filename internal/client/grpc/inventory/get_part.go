package inventory

import (
	"context"
	"order/internal/client/converter"
	repoModel "order/internal/repository/model"
	v1 "order/pkg/inventory"
)

func (c *Client) GetPart(ctx context.Context, partUUID string) (*repoModel.Part, error) {
	req, err := c.generatedClient.GetPart(ctx, &v1.GetPartRequest{
		Uuid: partUUID,
	})
	if err != nil {
		return nil, err
	}
	return converter.PartProtoToRepoModel(req.Part), nil
}

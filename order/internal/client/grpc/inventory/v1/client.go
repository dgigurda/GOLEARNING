package v1

import (
	"context"

	"google.golang.org/grpc"

	clientConverter "golearning/order/internal/client/converter"
	"golearning/order/internal/model"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

type Client struct {
	client inventoryv1.InventoryServiceClient
}

func NewClient(conn grpc.ClientConnInterface) *Client {
	return &Client{client: inventoryv1.NewInventoryServiceClient(conn)}
}

func NewFromGenerated(client inventoryv1.InventoryServiceClient) *Client {
	return &Client{client: client}
}

func (c *Client) ListParts(ctx context.Context, partUUIDs []string) ([]model.Part, error) {
	response, err := c.client.ListParts(ctx, &inventoryv1.ListPartsRequest{Filter: &inventoryv1.ListPartsFilter{Uuids: partUUIDs}})
	if err != nil {
		return nil, err
	}

	return clientConverter.PartsFromProto(response.GetParts()), nil
}

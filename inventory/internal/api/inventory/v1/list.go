package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inventoryConverter "golearning/inventory/internal/converter"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

func (a *API) ListParts(ctx context.Context, request *inventoryv1.ListPartsRequest) (*inventoryv1.ListPartsResponse, error) {
	parts, err := a.partService.ListParts(ctx, inventoryConverter.ProtoToListPartsFilter(request))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &inventoryv1.ListPartsResponse{Parts: inventoryConverter.PartsToProto(parts)}, nil
}

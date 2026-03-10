package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inventoryConverter "golearning/inventory/internal/converter"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

func (a *API) GetPart(ctx context.Context, request *inventoryv1.GetPartRequest) (*inventoryv1.GetPartResponse, error) {
	part, err := a.partService.GetPart(ctx, inventoryConverter.ProtoToGetPartRequest(request))
	if err != nil {
		if inventoryConverter.IsNotFoundError(err) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		if inventoryConverter.IsInvalidFilterError(err) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &inventoryv1.GetPartResponse{Part: inventoryConverter.PartToProto(part)}, nil
}

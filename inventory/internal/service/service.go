package service

import (
	"context"

	"golearning/inventory/internal/model"
)

type PartService interface {
	GetPart(ctx context.Context, filter model.GetPartResponse) (*model.Part, error)
	ListParts(ctx context.Context, filter model.ListPartsFilter) ([]*model.Part, error)
}

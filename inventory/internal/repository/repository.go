package repository

import (
	"context"

	repositoryModel "golearning/inventory/internal/repository/model"
)

type PartRepository interface {
	GetPart(ctx context.Context, filter repositoryModel.GetPartFilter) (*repositoryModel.Part, error)
	ListParts(ctx context.Context, filter repositoryModel.ListPartsFilter) ([]*repositoryModel.Part, error)
}

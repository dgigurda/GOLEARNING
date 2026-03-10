package part

import (
	"context"
	"sync"

	"golearning/inventory/internal/model"
	repositoryModel "golearning/inventory/internal/repository/model"
)

type Repository struct {
	mu     sync.RWMutex
	byID   map[int64]*repositoryModel.Part
	byUUID map[string]*repositoryModel.Part
}

func NewRepository() *Repository {
	seed := seedParts()
	byID := make(map[int64]*repositoryModel.Part, len(seed))
	byUUID := make(map[string]*repositoryModel.Part, len(seed))
	for _, part := range seed {
		byID[part.ID] = part
		byUUID[part.UUID] = part
	}

	return &Repository{
		byID:   byID,
		byUUID: byUUID,
	}
}

var _ interface {
	GetPart(context.Context, repositoryModel.GetPartFilter) (*repositoryModel.Part, error)
	ListParts(context.Context, repositoryModel.ListPartsFilter) ([]*repositoryModel.Part, error)
} = (*Repository)(nil)

func notFoundError() error {
	return model.ErrPartNotFound
}

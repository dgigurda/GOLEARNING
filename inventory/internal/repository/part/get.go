package part

import (
	"context"

	repositoryModel "golearning/inventory/internal/repository/model"
)

func (r *Repository) GetPart(_ context.Context, filter repositoryModel.GetPartFilter) (*repositoryModel.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if filter.UUID != "" {
		part, ok := r.byUUID[filter.UUID]
		if !ok {
			return nil, notFoundError()
		}
		return part, nil
	}

	if filter.ID != 0 {
		part, ok := r.byID[filter.ID]
		if !ok {
			return nil, notFoundError()
		}
		return part, nil
	}

	return nil, notFoundError()
}

package part

import (
	"context"
	"errors"

	"golearning/inventory/internal/model"
	repositoryConverter "golearning/inventory/internal/repository/converter"
)

func (s *Service) GetPart(ctx context.Context, filter model.GetPartResponse) (*model.Part, error) {
	if filter.ID == 0 && filter.UUID == "" {
		return nil, model.ErrInvalidGetPartFilter
	}

	repositoryPart, err := s.repository.GetPart(ctx, repositoryConverter.ToRepositoryGetPartFilter(filter))
	if err != nil {
		if errors.Is(err, model.ErrPartNotFound) {
			return nil, model.ErrPartNotFound
		}
		return nil, err
	}

	return repositoryConverter.ToModelPart(repositoryPart), nil
}

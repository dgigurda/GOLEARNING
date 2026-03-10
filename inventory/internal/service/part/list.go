package part

import (
	"context"

	"golearning/inventory/internal/model"
	repositoryConverter "golearning/inventory/internal/repository/converter"
)

func (s *Service) ListParts(ctx context.Context, filter model.ListPartsFilter) ([]*model.Part, error) {
	repositoryParts, err := s.repository.ListParts(ctx, repositoryConverter.ToRepositoryListPartsFilter(filter))
	if err != nil {
		return nil, err
	}

	parts := make([]*model.Part, 0, len(repositoryParts))
	for _, part := range repositoryParts {
		parts = append(parts, repositoryConverter.ToModelPart(part))
	}

	return parts, nil
}

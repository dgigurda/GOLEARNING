package part

import (
	"context"
	"slices"

	repositoryModel "golearning/inventory/internal/repository/model"
)

func (r *Repository) ListParts(_ context.Context, filter repositoryModel.ListPartsFilter) ([]*repositoryModel.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	parts := make([]*repositoryModel.Part, 0, len(r.byUUID))
	for _, part := range r.byUUID {
		if !matchesFilter(part, filter) {
			continue
		}
		parts = append(parts, part)
	}

	return parts, nil
}

func matchesFilter(part *repositoryModel.Part, filter repositoryModel.ListPartsFilter) bool {
	if len(filter.UUIDs) > 0 && !slices.Contains(filter.UUIDs, part.UUID) {
		return false
	}
	if len(filter.Names) > 0 && !slices.Contains(filter.Names, part.Name) {
		return false
	}
	if len(filter.Categories) > 0 && !slices.Contains(filter.Categories, part.Category) {
		return false
	}
	if len(filter.ManufacturerCountries) > 0 {
		if part.Manufacturer == nil || !slices.Contains(filter.ManufacturerCountries, part.Manufacturer.Country) {
			return false
		}
	}
	if len(filter.Tags) > 0 {
		hasTag := false
		for _, tag := range filter.Tags {
			if slices.Contains(part.Tags, tag) {
				hasTag = true
				break
			}
		}
		if !hasTag {
			return false
		}
	}

	return true
}

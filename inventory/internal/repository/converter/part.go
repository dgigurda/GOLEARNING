package converter

import (
	"golearning/inventory/internal/model"
	repositoryModel "golearning/inventory/internal/repository/model"
)

func ToRepositoryGetPartFilter(filter model.GetPartResponse) repositoryModel.GetPartFilter {
	return repositoryModel.GetPartFilter{
		ID:   filter.ID,
		UUID: filter.UUID,
	}
}

func ToRepositoryListPartsFilter(filter model.ListPartsFilter) repositoryModel.ListPartsFilter {
	categories := make([]repositoryModel.Category, 0, len(filter.Categories))
	for _, category := range filter.Categories {
		categories = append(categories, repositoryModel.Category(category))
	}

	return repositoryModel.ListPartsFilter{
		UUIDs:                 append([]string(nil), filter.UUIDs...),
		Names:                 append([]string(nil), filter.Names...),
		Categories:            categories,
		ManufacturerCountries: append([]string(nil), filter.ManufacturerCountries...),
		Tags:                  append([]string(nil), filter.Tags...),
	}
}

func ToModelPart(part *repositoryModel.Part) *model.Part {
	if part == nil {
		return nil
	}

	metadata := make(map[string]model.Metadata, len(part.Metadata))
	for key, value := range part.Metadata {
		metadata[key] = model.Metadata{
			ID:    value.ID,
			Value: toModelMetadataValue(value.Value),
		}
	}

	result := &model.Part{
		ID:            part.ID,
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      model.Category(part.Category),
		Tags:          append([]string(nil), part.Tags...),
		Metadata:      metadata,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}

	if part.Dimensions != nil {
		result.Dimensions = &model.Dimensions{
			ID:     part.Dimensions.ID,
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		}
	}

	if part.Manufacturer != nil {
		result.Manufacturer = &model.Manufacturer{
			ID:      part.Manufacturer.ID,
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		}
	}

	return result
}

func toModelMetadataValue(value repositoryModel.MetadataValue) model.MetadataValue {
	switch v := value.(type) {
	case repositoryModel.StringValue:
		return model.StringValue{V: v.V}
	case repositoryModel.Int64Value:
		return model.Int64Value{V: v.V}
	case repositoryModel.DoubleValue:
		return model.DoubleValue{V: v.V}
	case repositoryModel.BoolValue:
		return model.BoolValue{V: v.V}
	default:
		return nil
	}
}

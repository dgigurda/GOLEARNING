package converter

import (
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"golearning/inventory/internal/model"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

func PartToModel(part *inventoryv1.Part) model.Part {
	if part == nil {
		return model.Part{}
	}
	var UpdateAt *time.Time
	if part.UpdateAt != nil {
		tmp := part.UpdateAt.AsTime()
		UpdateAt = &tmp
	}
	var CreatedAt *time.Time
	if part.CreatedAt != nil {
		tmp := part.CreatedAt.AsTime()
		CreatedAt = &tmp
	}
	result := model.Part{
		UUID:          part.Uuid,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      model.Category(part.Category.Number()),
		Dimensions: &model.Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		},
		Manufacturer: &model.Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		},
		Tags:      append([]string(nil), part.Tags...),
		Metadata:  MetadataToModel(part.Metadata),
		CreatedAt: *CreatedAt,
		UpdatedAt: *UpdateAt,
	}
	return result
}

func MetadataToModel(metadata map[string]*inventoryv1.Metadata) map[string]model.Metadata {
	resultMap := make(map[string]model.Metadata)
	for k, v := range metadata {
		if v == nil {
			continue
		}
		item := model.Metadata{
			ID: v.GetId(),
		}
		switch value := v.GetValue().(type) {
		case *inventoryv1.Metadata_StringValue:
			item.Value = model.StringValue{V: value.StringValue}
		case *inventoryv1.Metadata_Int64Value:
			item.Value = model.Int64Value{V: value.Int64Value}
		case *inventoryv1.Metadata_DoubleValue:
			item.Value = model.DoubleValue{V: value.DoubleValue}
		case *inventoryv1.Metadata_BoolValue:
			item.Value = model.BoolValue{V: value.BoolValue}
		default:
			item.Value = nil
		}
		resultMap[k] = item
	}

	return resultMap
}

func PartToProto(part *model.Part) *inventoryv1.Part {
	if part == nil {
		return nil
	}

	metadata := make(map[string]*inventoryv1.Metadata, len(part.Metadata))
	for key, value := range part.Metadata {
		metadata[key] = metadataToProto(value)
	}

	result := &inventoryv1.Part{
		Id:            part.ID,
		Uuid:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      categoryToProto(part.Category),
		Tags:          append([]string(nil), part.Tags...),
		Metadata:      metadata,
		CreatedAt:     timestamppb.New(part.CreatedAt),
		UpdateAt:      timestamppb.New(part.UpdatedAt),
	}

	if part.Dimensions != nil {
		result.Dimensions = &inventoryv1.Dimensions{
			Id:     part.Dimensions.ID,
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		}
	}

	if part.Manufacturer != nil {
		result.Manufacturer = &inventoryv1.Manufacturer{
			Id:      part.Manufacturer.ID,
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		}
	}

	return result
}

func ProtoToGetPartRequest(request *inventoryv1.GetPartRequest) model.GetPartResponse {
	if request == nil {
		return model.GetPartResponse{}
	}

	return model.GetPartResponse{
		ID:   request.GetId(),
		UUID: request.GetUuid(),
	}
}

func ProtoToListPartsFilter(request *inventoryv1.ListPartsRequest) model.ListPartsFilter {
	if request == nil || request.GetFilter() == nil {
		return model.ListPartsFilter{}
	}

	filter := request.GetFilter()
	categories := make([]model.Category, 0, len(filter.GetCategorys()))
	for _, category := range filter.GetCategorys() {
		categories = append(categories, categoryFromProto(category))
	}

	return model.ListPartsFilter{
		UUIDs:                 append([]string(nil), filter.GetUuids()...),
		Names:                 append([]string(nil), filter.GetNames()...),
		Categories:            categories,
		ManufacturerCountries: append([]string(nil), filter.GetManufacturerCountries()...),
		Tags:                  append([]string(nil), filter.GetTags()...),
	}
}

func metadataToProto(metadata model.Metadata) *inventoryv1.Metadata {
	result := &inventoryv1.Metadata{Id: metadata.ID}

	switch value := metadata.Value.(type) {
	case model.StringValue:
		result.Value = &inventoryv1.Metadata_StringValue{StringValue: value.V}
	case model.Int64Value:
		result.Value = &inventoryv1.Metadata_Int64Value{Int64Value: value.V}
	case model.DoubleValue:
		result.Value = &inventoryv1.Metadata_DoubleValue{DoubleValue: value.V}
	case model.BoolValue:
		result.Value = &inventoryv1.Metadata_BoolValue{BoolValue: value.V}
	default:
		result.Value = nil
	}

	return result
}

func categoryToProto(category model.Category) inventoryv1.Category {
	switch category {
	case model.CategoryEngine:
		return inventoryv1.Category_CATEGORY_ENGINE
	case model.CategoryFuel:
		return inventoryv1.Category_CATEGORY_FUEL
	case model.CategoryPorthole:
		return inventoryv1.Category_CATEGORY_PORTHOLE
	case model.CategoryWing:
		return inventoryv1.Category_CATEGORY_WING
	default:
		return inventoryv1.Category_CATEGORY_UNSPECIFIED
	}
}

func categoryFromProto(category inventoryv1.Category) model.Category {
	switch category {
	case inventoryv1.Category_CATEGORY_ENGINE:
		return model.CategoryEngine
	case inventoryv1.Category_CATEGORY_FUEL:
		return model.CategoryFuel
	case inventoryv1.Category_CATEGORY_PORTHOLE:
		return model.CategoryPorthole
	case inventoryv1.Category_CATEGORY_WING:
		return model.CategoryWing
	default:
		return model.CategoryUnspecified
	}
}

func PartsToProto(parts []*model.Part) []*inventoryv1.Part {
	result := make([]*inventoryv1.Part, 0, len(parts))
	for _, part := range parts {
		result = append(result, PartToProto(part))
	}
	return result
}

func IsNotFoundError(err error) bool {
	return errors.Is(err, model.ErrPartNotFound)
}

func IsInvalidFilterError(err error) bool {
	return errors.Is(err, model.ErrInvalidGetPartFilter)
}

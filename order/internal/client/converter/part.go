package converter

import (
	"golearning/order/internal/model"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

func PartsFromProto(parts []*inventoryv1.Part) []model.Part {
	result := make([]model.Part, 0, len(parts))
	for _, part := range parts {
		if part == nil {
			continue
		}
		result = append(result, model.Part{UUID: part.GetUuid(), Price: float32(part.GetPrice())})
	}
	return result
}

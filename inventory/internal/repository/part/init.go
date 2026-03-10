package part

import (
	"time"

	repositoryModel "golearning/inventory/internal/repository/model"
)

func seedParts() []*repositoryModel.Part {
	now := time.Now().UTC()

	return []*repositoryModel.Part{
		{
			ID:            1,
			UUID:          "11111111-1111-1111-1111-111111111111",
			Name:          "Main Engine X1",
			Description:   "Primary engine for medium spacecraft",
			Price:         150000,
			StockQuantity: 8,
			Category:      repositoryModel.CategoryEngine,
			Dimensions: &repositoryModel.Dimensions{
				ID:     1,
				Length: 4.2,
				Width:  2.1,
				Height: 2.4,
				Weight: 980,
			},
			Manufacturer: &repositoryModel.Manufacturer{
				ID:      1,
				Name:    "Orbital Dynamics",
				Country: "Germany",
				Website: "https://orbital.example",
			},
			Tags: []string{"engine", "main", "x1"},
			Metadata: map[string]repositoryModel.Metadata{
				"is_certified": {
					ID:    1,
					Value: repositoryModel.BoolValue{V: true},
				},
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:            2,
			UUID:          "22222222-2222-2222-2222-222222222222",
			Name:          "Fuel Tank F9",
			Description:   "Composite fuel tank",
			Price:         42000,
			StockQuantity: 20,
			Category:      repositoryModel.CategoryFuel,
			Dimensions: &repositoryModel.Dimensions{
				ID:     2,
				Length: 2.8,
				Width:  1.9,
				Height: 1.9,
				Weight: 420,
			},
			Manufacturer: &repositoryModel.Manufacturer{
				ID:      2,
				Name:    "Nova Fuel Systems",
				Country: "USA",
				Website: "https://nova.example",
			},
			Tags: []string{"fuel", "tank"},
			Metadata: map[string]repositoryModel.Metadata{
				"batch": {
					ID:    2,
					Value: repositoryModel.StringValue{V: "F9-2026-01"},
				},
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

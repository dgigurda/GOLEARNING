package v1

import (
	inventoryConverter "golearning/inventory/internal/converter"
	"golearning/inventory/internal/service"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

type API struct {
	inventoryv1.UnimplementedInventoryServiceServer

	partService service.PartService
}

func NewAPI(partService service.PartService) *API {
	return &API{partService: partService}
}

var (
	_ inventoryv1.InventoryServiceServer = (*API)(nil)
	_                                    = inventoryConverter.IsNotFoundError
)

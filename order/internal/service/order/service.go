package order

import (
	"context"

	"github.com/google/uuid"

	"golearning/order/internal/model"
	"golearning/order/internal/repository"
	servicepkg "golearning/order/internal/service"
)

type InventoryClient interface {
	ListParts(ctx context.Context, partUUIDs []string) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUUID, userUUID, paymentMethod string) (string, error)
}

type Service struct {
	repository      repository.OrderRepository
	inventoryClient InventoryClient
	paymentClient   PaymentClient
	newUUID         func() string
}

func NewService(repository repository.OrderRepository, inventoryClient InventoryClient, paymentClient PaymentClient) servicepkg.OrderService {
	return &Service{
		repository:      repository,
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
		newUUID:         uuid.NewString,
	}
}

package order

import (
	"context"
	"errors"

	"golearning/order/internal/model"
)

func (s *Service) Create(ctx context.Context, request model.CreateOrderRequest) (model.CreateOrderResponse, error) {
	if request.UserUUID == "" || len(request.PartUUIDs) == 0 {
		return model.CreateOrderResponse{}, model.ErrInvalidCreateOrderRequest
	}

	parts, err := s.inventoryClient.ListParts(ctx, request.PartUUIDs)
	if err != nil {
		return model.CreateOrderResponse{}, errors.Join(model.ErrExternalInventory, err)
	}
	if len(parts) != len(request.PartUUIDs) {
		return model.CreateOrderResponse{}, model.ErrPartsNotFound
	}

	var totalPrice float32
	for _, part := range parts {
		totalPrice += part.Price
	}

	createdOrder := model.Order{
		OrderUUID:  s.newUUID(),
		UserUUID:   request.UserUUID,
		PartUUIDs:  append([]string(nil), request.PartUUIDs...),
		TotalPrice: totalPrice,
		Status:     model.OrderStatusPendingPayment,
	}

	if err := s.repository.Create(ctx, createdOrder); err != nil {
		return model.CreateOrderResponse{}, err
	}

	return model.CreateOrderResponse{OrderUUID: createdOrder.OrderUUID, TotalPrice: totalPrice}, nil
}

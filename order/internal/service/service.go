package service

import (
	"context"

	"golearning/order/internal/model"
)

type OrderService interface {
	Create(ctx context.Context, request model.CreateOrderRequest) (model.CreateOrderResponse, error)
	Get(ctx context.Context, uuid string) (model.Order, error)
	Pay(ctx context.Context, request model.PayOrderRequest) (model.PayOrderResponse, error)
	Cancel(ctx context.Context, uuid string) error
}

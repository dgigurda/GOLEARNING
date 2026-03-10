package service

import (
	"context"

	"golearning/payment/internal/model"
)

type PaymentService interface {
	Pay(ctx context.Context, request model.PayOrderRequest) (model.PayOrderResponse, error)
}

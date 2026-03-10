package repository

import (
	"context"

	"golearning/payment/internal/model"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment model.Payment) error
}

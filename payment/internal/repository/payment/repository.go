package payment

import (
	"context"
	"sync"

	"golearning/payment/internal/model"
)

type Repository struct {
	mu       sync.Mutex
	payments map[string]model.Payment
}

func NewRepository() *Repository {
	return &Repository{payments: make(map[string]model.Payment)}
}

func (r *Repository) Save(_ context.Context, payment model.Payment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.payments[payment.TransactionUUID] = payment
	return nil
}

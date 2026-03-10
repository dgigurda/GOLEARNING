package order

import (
	"context"
	"sync"

	"golearning/order/internal/model"
	repositoryConverter "golearning/order/internal/repository/converter"
	repositoryModel "golearning/order/internal/repository/model"
)

type Repository struct {
	mu     sync.RWMutex
	orders map[string]repositoryModel.Order
}

func NewRepository() *Repository {
	return &Repository{orders: make(map[string]repositoryModel.Order)}
}

func (r *Repository) Get(_ context.Context, uuid string) (model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.orders[uuid]
	if !ok {
		return model.Order{}, model.ErrOrderNotFound
	}

	return repositoryConverter.ToModelOrder(order), nil
}

func (r *Repository) Create(_ context.Context, order model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.OrderUUID] = repositoryConverter.ToRepositoryOrder(order)
	return nil
}

func (r *Repository) Update(_ context.Context, order model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.orders[order.OrderUUID]; !ok {
		return model.ErrOrderNotFound
	}

	r.orders[order.OrderUUID] = repositoryConverter.ToRepositoryOrder(order)
	return nil
}

var _ interface {
	Get(context.Context, string) (model.Order, error)
	Create(context.Context, model.Order) error
	Update(context.Context, model.Order) error
} = (*Repository)(nil)

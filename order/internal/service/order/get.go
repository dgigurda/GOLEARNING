package order

import (
	"context"

	"golearning/order/internal/model"
)

func (s *Service) Get(ctx context.Context, uuid string) (model.Order, error) {
	if uuid == "" {
		return model.Order{}, model.ErrOrderNotFound
	}

	return s.repository.Get(ctx, uuid)
}

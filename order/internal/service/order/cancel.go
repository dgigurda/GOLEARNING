package order

import (
	"context"

	"golearning/order/internal/model"
)

func (s *Service) Cancel(ctx context.Context, uuid string) error {
	if uuid == "" {
		return model.ErrOrderNotFound
	}

	order, err := s.repository.Get(ctx, uuid)
	if err != nil {
		return err
	}
	if order.Status != model.OrderStatusPendingPayment {
		return model.ErrOrderConflict
	}

	order.Status = model.OrderStatusCancelled
	return s.repository.Update(ctx, order)
}

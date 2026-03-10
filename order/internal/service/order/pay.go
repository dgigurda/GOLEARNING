package order

import (
	"context"
	"errors"

	"golearning/order/internal/model"
)

func (s *Service) Pay(ctx context.Context, request model.PayOrderRequest) (model.PayOrderResponse, error) {
	if request.OrderUUID == "" || request.PaymentMethod == "" {
		return model.PayOrderResponse{}, model.ErrInvalidPayOrderRequest
	}

	order, err := s.repository.Get(ctx, request.OrderUUID)
	if err != nil {
		return model.PayOrderResponse{}, err
	}
	if order.Status != model.OrderStatusPendingPayment {
		return model.PayOrderResponse{}, model.ErrOrderConflict
	}

	transactionUUID, err := s.paymentClient.PayOrder(ctx, order.OrderUUID, order.UserUUID, request.PaymentMethod)
	if err != nil {
		return model.PayOrderResponse{}, errors.Join(model.ErrExternalPayment, err)
	}

	order.Status = model.OrderStatusPaid
	order.TransactionUUID = transactionUUID
	order.PaymentMethod = request.PaymentMethod

	if err := s.repository.Update(ctx, order); err != nil {
		return model.PayOrderResponse{}, err
	}

	return model.PayOrderResponse{TransactionUUID: transactionUUID}, nil
}

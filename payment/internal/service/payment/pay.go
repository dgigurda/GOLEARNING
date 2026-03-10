package payment

import (
	"context"

	"github.com/google/uuid"

	"golearning/payment/internal/model"
)

func (s *Service) Pay(ctx context.Context, request model.PayOrderRequest) (model.PayOrderResponse, error) {
	if uuid.Validate(request.OrderUUID) != nil {
		return model.PayOrderResponse{}, model.ErrInvalidOrderUUID
	}
	if uuid.Validate(request.UserUUID) != nil {
		return model.PayOrderResponse{}, model.ErrInvalidUserUUID
	}
	if request.PaymentMethod == 0 {
		return model.PayOrderResponse{}, model.ErrInvalidPaymentMethod
	}

	transactionUUID := s.newUUID()
	payment := model.Payment{
		OrderUUID:       request.OrderUUID,
		UserUUID:        request.UserUUID,
		PaymentMethod:   request.PaymentMethod,
		TransactionUUID: transactionUUID,
	}

	if err := s.repository.Save(ctx, payment); err != nil {
		return model.PayOrderResponse{}, err
	}

	return model.PayOrderResponse{TransactionUUID: transactionUUID}, nil
}

package order

import (
	"context"

	"golearning/order/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

func (s *ServiceSuite) TestPayInvalidUUID() {
	var (
		payOrderRequest = model.PayOrderRequest{
			OrderUUID: "",
		}
	)

	res, err := s.service.Pay(s.ctx, payOrderRequest)
	s.Require().ErrorIs(err, model.ErrInvalidPayOrderRequest)
	s.Require().Equal(model.PayOrderResponse{}, res)
}

func (s *ServiceSuite) TestPayRepoErr() {
	var (
		repoErr       = gofakeit.Error()
		orderUUID     = gofakeit.UUID()
		paymentMethod = "CARD"

		payOrderRequest = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			PaymentMethod: paymentMethod,
		}
	)

	s.repo.On("Get", s.ctx, payOrderRequest.OrderUUID).Return(model.Order{}, repoErr)

	res, err := s.service.Pay(s.ctx, payOrderRequest)
	s.Require().ErrorIs(err, repoErr)
	s.Require().Equal(model.PayOrderResponse{}, res)
}

func (s *ServiceSuite) TestPayOrderStatusErr() {
	var (
		orderUUID     = gofakeit.UUID()
		paymentMethod = "CARD"
		orderStatus   = model.OrderStatusCancelled

		payOrderRequest = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			PaymentMethod: paymentMethod,
		}

		order = model.Order{
			Status: orderStatus,
		}
	)

	s.repo.
		On("Get", s.ctx, payOrderRequest.OrderUUID).
		Return(order, nil)

	res, err := s.service.Pay(s.ctx, payOrderRequest)
	s.Require().ErrorIs(err, model.ErrOrderConflict)
	s.Require().Equal(model.PayOrderResponse{}, res)
}

func (s *ServiceSuite) TestPayOrderPayingErr() {
	var (
		payErr        = gofakeit.Error()
		userUUID      = gofakeit.UUID()
		orderUUID     = gofakeit.UUID()
		paymentMethod = "CARD"
		orderStatus   = model.OrderStatusPendingPayment

		payOrderRequest = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			UserUUID:      userUUID,
			PaymentMethod: paymentMethod,
		}

		order = model.Order{
			Status: orderStatus,
		}
	)

	s.repo.
		On("Get", s.ctx, payOrderRequest.OrderUUID).
		Return(order, nil)

	s.paymentClient.payOrderFn = func(ctx context.Context, orderUUID string, userUUID string, paymentMethod string) (string, error) {
		return "", payErr
	}

	res, err := s.service.Pay(s.ctx, payOrderRequest)
	s.Require().ErrorIs(err, payErr)
	s.Require().Equal(model.PayOrderResponse{}, res)
}

func (s *ServiceSuite) TestPayOrderUpdateErr() {
	var (
		updateErr       = gofakeit.Error()
		transactionUUID = gofakeit.UUID()
		userUUID        = gofakeit.UUID()
		orderUUID       = gofakeit.UUID()
		paymentMethod   = "CARD"
		orderStatus     = model.OrderStatusPendingPayment

		payOrderRequest = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			UserUUID:      userUUID,
			PaymentMethod: paymentMethod,
		}

		order = model.Order{
			Status:          orderStatus,
			TransactionUUID: transactionUUID,
			PaymentMethod:   paymentMethod,
		}
	)

	s.repo.
		On("Get", s.ctx, payOrderRequest.OrderUUID).
		Return(order, nil)

	s.paymentClient.payOrderFn = func(ctx context.Context, orderUUID string, userUUID string, paymentMethod string) (string, error) {
		return transactionUUID, nil
	}

	order.Status = model.OrderStatusPaid

	s.repo.
		On("Update", s.ctx, order).Return(updateErr)

	res, err := s.service.Pay(s.ctx, payOrderRequest)
	s.Require().ErrorIs(err, updateErr)
	s.Require().Equal(model.PayOrderResponse{}, res)
}

func (s *ServiceSuite) TestPayOrderSuccess() {
	var (
		transactionUUID = gofakeit.UUID()
		userUUID        = gofakeit.UUID()
		orderUUID       = gofakeit.UUID()
		paymentMethod   = "CARD"
		orderStatus     = model.OrderStatusPendingPayment

		payOrderRequest = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			UserUUID:      userUUID,
			PaymentMethod: paymentMethod,
		}

		order = model.Order{
			Status:          orderStatus,
			TransactionUUID: transactionUUID,
			PaymentMethod:   paymentMethod,
		}
	)

	s.repo.
		On("Get", s.ctx, payOrderRequest.OrderUUID).
		Return(order, nil)

	s.paymentClient.payOrderFn = func(ctx context.Context, orderUUID string, userUUID string, paymentMethod string) (string, error) {
		return transactionUUID, nil
	}

	order.Status = model.OrderStatusPaid

	s.repo.
		On("Update", s.ctx, order).Return(nil)

	res, err := s.service.Pay(s.ctx, payOrderRequest)
	s.Require().ErrorIs(err, nil)
	s.Require().Equal(model.PayOrderResponse{TransactionUUID: transactionUUID}, res)
}

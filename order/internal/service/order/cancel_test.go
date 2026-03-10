package order

import (
	"errors"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/mock"

	"golearning/order/internal/model"
)

func (s *ServiceSuite) TestCancelInvalidUUID() {
	err := s.service.Cancel(s.ctx, "")
	s.Require().ErrorIs(err, model.ErrOrderNotFound)
}

func (s *ServiceSuite) TestCancelGetError() {
	orderUUID := gofakeit.UUID()
	repoErr := errors.New("db get failed")

	s.repo.On("Get", s.ctx, orderUUID).
		Return(model.Order{}, repoErr)

	err := s.service.Cancel(s.ctx, orderUUID)
	s.Require().ErrorIs(err, repoErr)
}

func (s *ServiceSuite) TestCancelConflict() {
	orderUUID := gofakeit.UUID()
	order := model.Order{
		OrderUUID: orderUUID,
		Status:    model.OrderStatusPaid,
	}

	s.repo.On("Get", s.ctx, orderUUID).
		Return(order, nil)

	err := s.service.Cancel(s.ctx, orderUUID)
	s.Require().ErrorIs(err, model.ErrOrderConflict)
}

func (s *ServiceSuite) TestCancelSuccess() {
	orderUUID := gofakeit.UUID()
	order := model.Order{
		OrderUUID: orderUUID,
		Status:    model.OrderStatusPendingPayment,
	}

	s.repo.On("Get", s.ctx, orderUUID).
		Return(order, nil)

	s.repo.On("Update", s.ctx, mock.MatchedBy(func(o model.Order) bool {
		return o.OrderUUID == orderUUID && o.Status == model.OrderStatusCancelled
	})).Return(nil)

	err := s.service.Cancel(s.ctx, orderUUID)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestCancelUpdateError() {
	orderUUID := gofakeit.UUID()
	repoErr := errors.New("db update failed")
	order := model.Order{
		OrderUUID: orderUUID,
		Status:    model.OrderStatusPendingPayment,
	}

	s.repo.On("Get", s.ctx, orderUUID).
		Return(order, nil)

	s.repo.On("Update", s.ctx, mock.MatchedBy(func(o model.Order) bool {
		return o.OrderUUID == orderUUID && o.Status == model.OrderStatusCancelled
	})).Return(repoErr)

	err := s.service.Cancel(s.ctx, orderUUID)
	s.Require().ErrorIs(err, repoErr)
}

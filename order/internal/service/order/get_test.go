package order

import (
	"github.com/brianvoe/gofakeit/v6"

	"golearning/order/internal/model"
)

func (s *ServiceSuite) TestGetInvalidUUID() {

	res, err := s.service.Get(s.ctx, "")
	s.Require().ErrorIs(err, model.ErrOrderNotFound)
	s.Require().Equal(model.Order{}, res)
}

func (s *ServiceSuite) TestGetRepoError() {
	repoErr := gofakeit.Error()
	uuid := gofakeit.UUID()

	s.repo.On("Get", s.ctx, uuid).Return(model.Order{}, repoErr)

	res, err := s.service.Get(s.ctx, uuid)
	s.Require().ErrorIs(err, repoErr)
	s.Require().Equal(model.Order{}, res)
}
func (s *ServiceSuite) TestGetSuccess() {
	uuid := gofakeit.UUID()
	orderUUID := gofakeit.UUID()

	s.repo.On("Get", s.ctx, uuid).Return(model.Order{OrderUUID: orderUUID}, nil)

	res, err := s.service.Get(s.ctx, uuid)
	s.Require().ErrorIs(err, nil)
	s.Require().Equal(orderUUID, res.OrderUUID)
}

package order

import (
	"context"
	"errors"

	"github.com/brianvoe/gofakeit/v6"

	"golearning/order/internal/model"
)

func (s *ServiceSuite) TestCreateInvalidUUID() {
	req := model.CreateOrderRequest{
		UserUUID: "",
	}

	res, err := s.service.Create(s.ctx, req)
	s.Require().ErrorIs(err, model.ErrInvalidCreateOrderRequest)
	s.Require().Equal(model.CreateOrderResponse{}, res)
}

func (s *ServiceSuite) TestCreateListPartsError() {
	var (
		userUUID  = gofakeit.UUID()
		partUUIDs = []string{gofakeit.UUID()}

		req = model.CreateOrderRequest{
			UserUUID:  userUUID,
			PartUUIDs: partUUIDs,
		}
	)

	inventoryErr := errors.New("inventory down")
	s.inventoryClient.listPartsFn = func(ctx context.Context, ids []string) ([]model.Part, error) {
		s.Require().Equal(partUUIDs, ids)
		return nil, inventoryErr
	}

	res, err := s.service.Create(s.ctx, req)
	s.Require().ErrorIs(err, model.ErrExternalInventory)
	s.Require().ErrorIs(err, inventoryErr)
	s.Require().Equal(model.CreateOrderResponse{}, res)
}

func (s *ServiceSuite) TestCreateLenError() {
	var (
		userUUID  = gofakeit.UUID()
		partUUIDs = []string{gofakeit.UUID()}

		req = model.CreateOrderRequest{
			UserUUID:  userUUID,
			PartUUIDs: partUUIDs,
		}
	)

	lenErr := model.ErrPartsNotFound
	s.inventoryClient.listPartsFn = func(ctx context.Context, ids []string) ([]model.Part, error) {
		s.Require().Equal(partUUIDs, ids)
		return []model.Part{}, nil
	}

	res, err := s.service.Create(s.ctx, req)
	s.Require().ErrorIs(err, model.ErrPartsNotFound)
	s.Require().ErrorIs(err, lenErr)
	s.Require().Equal(model.CreateOrderResponse{}, res)
}

func (s *ServiceSuite) TestCreateError() {
	var (
		repoErr = gofakeit.Error()

		orderUUID = gofakeit.UUID()
		userUUID  = gofakeit.UUID()
		partUUIDs = []string{gofakeit.UUID()}

		req = model.CreateOrderRequest{
			UserUUID:  userUUID,
			PartUUIDs: partUUIDs,
		}

		order = model.Order{
			OrderUUID:  orderUUID,
			UserUUID:   userUUID,
			PartUUIDs:  partUUIDs,
			TotalPrice: 100,
			Status:     model.OrderStatusPendingPayment,
		}
	)

	s.service = &Service{
		repository:      s.repo,
		inventoryClient: s.inventoryClient,
		paymentClient:   s.paymentClient,
		newUUID:         func() string { return orderUUID },
	}

	s.inventoryClient.listPartsFn = func(ctx context.Context, ids []string) ([]model.Part, error) {
		s.Require().Equal(partUUIDs, ids)
		return []model.Part{
			{UUID: partUUIDs[0], Price: 100},
		}, nil
	}
	s.repo.On("Create", s.ctx, order).Return(repoErr)

	res, err := s.service.Create(s.ctx, req)
	s.Require().ErrorIs(err, repoErr)
	s.Require().Equal(model.CreateOrderResponse{}, res)
}

func (s *ServiceSuite) TestCreateSuccess() {
	var (
		orderUUID = gofakeit.UUID()
		userUUID  = gofakeit.UUID()
		partUUIDs = []string{gofakeit.UUID()}

		req = model.CreateOrderRequest{
			UserUUID:  userUUID,
			PartUUIDs: partUUIDs,
		}

		order = model.Order{
			OrderUUID:  orderUUID,
			UserUUID:   userUUID,
			PartUUIDs:  partUUIDs,
			TotalPrice: 100,
			Status:     model.OrderStatusPendingPayment,
		}
	)

	s.service = &Service{
		repository:      s.repo,
		inventoryClient: s.inventoryClient,
		paymentClient:   s.paymentClient,
		newUUID:         func() string { return orderUUID },
	}

	s.inventoryClient.listPartsFn = func(ctx context.Context, ids []string) ([]model.Part, error) {
		s.Require().Equal(partUUIDs, ids)
		return []model.Part{
			{UUID: partUUIDs[0], Price: 100},
		}, nil
	}
	s.repo.On("Create", s.ctx, order).Return(nil)

	res, err := s.service.Create(s.ctx, req)
	s.Require().ErrorIs(err, nil)
	s.Require().Equal(orderUUID, res.OrderUUID)
	s.Require().Equal(order.TotalPrice, res.TotalPrice)
}

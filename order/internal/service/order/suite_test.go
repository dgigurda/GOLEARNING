package order

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"golearning/order/internal/model"
	mocks "golearning/order/internal/repository/mocks"
	servicepkg "golearning/order/internal/service"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repo *mocks.OrderRepository

	inventoryClient *inventoryClientStub
	paymentClient   *paymentClientStub

	service servicepkg.OrderService
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.repo = mocks.NewOrderRepository(s.T())
	s.inventoryClient = &inventoryClientStub{}
	s.paymentClient = &paymentClientStub{}

	s.service = NewService(s.repo, s.inventoryClient, s.paymentClient)
}

func TestOrderServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type inventoryClientStub struct {
	listPartsFn func(ctx context.Context, partUUIDs []string) ([]model.Part, error)
}

func (c *inventoryClientStub) ListParts(ctx context.Context, partUUIDs []string) ([]model.Part, error) {
	if c.listPartsFn == nil {
		return nil, nil
	}
	return c.listPartsFn(ctx, partUUIDs)
}

type paymentClientStub struct {
	payOrderFn func(ctx context.Context, orderUUID, userUUID, paymentMethod string) (string, error)
}

func (c *paymentClientStub) PayOrder(ctx context.Context, orderUUID, userUUID, paymentMethod string) (string, error) {
	if c.payOrderFn == nil {
		return "", nil
	}
	return c.payOrderFn(ctx, orderUUID, userUUID, paymentMethod)
}

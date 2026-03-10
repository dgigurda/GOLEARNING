package payment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	mocks "golearning/payment/internal/repository/mocks"
	servicepkg "golearning/payment/internal/service"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repo *mocks.PaymentRepository

	service servicepkg.PaymentService // depend on interface
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.repo = mocks.NewPaymentRepository(s.T())
	s.service = NewService(s.repo) // constructor returns interface
}

func TestPartServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

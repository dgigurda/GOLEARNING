package part

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	mocks "golearning/inventory/internal/repository/mocks"
	servicepkg "golearning/inventory/internal/service"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repo *mocks.PartRepository

	service servicepkg.PartService // depend on interface
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.repo = mocks.NewPartRepository(s.T())
	s.service = NewService(s.repo) // constructor returns interface
}

func TestPartServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

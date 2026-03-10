package v1

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golearning/inventory/internal/model"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

func (s *APISuite) TestGetSuccess() {
	uuid := gofakeit.UUID()
	req := &inventoryv1.GetPartRequest{Uuid: uuid}
	expectedResponse := model.GetPartResponse{
		UUID: uuid,
	}

	expectedPart := &model.Part{
		UUID: uuid,
	}

	s.inventoryService.
		On("GetPart", s.ctx, expectedResponse).
		Return(expectedPart, nil)

	res, err := s.api.GetPart(s.ctx, req)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.NotNil(s.T(), res.Part)
	require.Equal(s.T(), uuid, res.Part.Uuid)
}

func (s *APISuite) TestGetServiceError() {
	serviceErr := gofakeit.Error()
	uuid := gofakeit.UUID()
	request := &inventoryv1.GetPartRequest{
		Uuid: uuid,
	}
	expectedResponse := model.GetPartResponse{
		UUID: uuid,
	}
	s.inventoryService.
		On("GetPart", s.ctx, expectedResponse).
		Return((*model.Part)(nil), serviceErr)

	res, err := s.api.GetPart(s.ctx, request)
	s.Require().Error(err)
	s.Require().Equal(codes.Internal, status.Code(err))
	s.Require().Nil(res)
}

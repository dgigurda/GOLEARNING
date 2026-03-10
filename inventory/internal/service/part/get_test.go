package part

import (
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appModel "golearning/inventory/internal/model"
	repoModel "golearning/inventory/internal/repository/model"
)

func (s *ServiceSuite) TestGetSuccess() {
	uuid := gofakeit.UUID()

	req := appModel.GetPartResponse{UUID: uuid}

	expectedRepoFilter := repoModel.GetPartFilter{UUID: uuid}
	repoPart := &repoModel.Part{UUID: uuid}

	s.repo.
		On("GetPart", s.ctx, expectedRepoFilter).
		Return(repoPart, nil)

	res, err := s.service.GetPart(s.ctx, req)
	s.Require().NoError(err)
	s.Require().Equal(uuid, res.UUID)
}

func (s *ServiceSuite) TestGetServiceError() {
	serviceErr := gofakeit.Error()
	uuid := gofakeit.UUID()

	req := appModel.GetPartResponse{UUID: uuid}

	expectedRequest := repoModel.GetPartFilter{
		UUID: uuid,
	}
	s.repo.
		On("GetPart", s.ctx, expectedRequest).
		Return((*repoModel.Part)(nil), serviceErr)

	res, err := s.service.GetPart(s.ctx, req)
	s.Require().Error(err)
	s.Require().Equal(codes.Unknown, status.Code(err))
	s.Require().ErrorIs(err, serviceErr)
	s.Require().Nil(res)

}

func (s *ServiceSuite) TestGetLenError() {
	serviceErr := appModel.ErrInvalidGetPartFilter

	req := appModel.GetPartResponse{UUID: ""}

	res, err := s.service.GetPart(s.ctx, req)
	s.Require().Error(err)
	s.Require().ErrorIs(err, serviceErr)
	s.Require().Nil(res)

}

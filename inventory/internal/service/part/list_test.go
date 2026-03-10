package part

import (
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appModel "golearning/inventory/internal/model"
	repoModel "golearning/inventory/internal/repository/model"
)

func (s *ServiceSuite) TestListSuccess() {
	var (
		uuid    = gofakeit.UUID()
		name    = gofakeit.Word()
		country = gofakeit.Country()
		tag     = gofakeit.Word()

		req = appModel.ListPartsFilter{
			UUIDs:                 []string{uuid},
			Names:                 []string{name},
			Categories:            []appModel.Category{appModel.CategoryEngine},
			ManufacturerCountries: []string{country},
			Tags:                  []string{tag},
		}

		expectedFilter = repoModel.ListPartsFilter{
			UUIDs:                 []string{uuid},
			Names:                 []string{name},
			Categories:            []repoModel.Category{repoModel.CategoryEngine},
			ManufacturerCountries: []string{country},
			Tags:                  []string{tag},
		}

		expectedParts = []*repoModel.Part{
			{
				UUID: uuid,
				Name: name,
			},
		}
	)

	s.repo.
		On("ListParts", s.ctx, expectedFilter).
		Return(expectedParts, nil)

	res, err := s.service.ListParts(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Len(res, 1)
	s.Require().Equal(uuid, res[0].UUID)
	s.Require().Equal(name, res[0].Name)
}

func (s *ServiceSuite) TestListServiceError() {
	var (
		serviceErr = gofakeit.Error()

		req = appModel.ListPartsFilter{
			UUIDs: []string{gofakeit.UUID()},
		}

		expectedFilter = repoModel.ListPartsFilter{
			UUIDs:      req.UUIDs,
			Categories: []repoModel.Category{},
		}
	)

	s.repo.On("ListParts", s.ctx, expectedFilter).
		Return(([]*repoModel.Part)(nil), serviceErr)

	res, err := s.service.ListParts(s.ctx, req)
	s.Require().Error(err)
	s.Require().Equal(codes.Unknown, status.Code(err))
	s.Require().ErrorIs(err, serviceErr)
	s.Require().Nil(res)
}

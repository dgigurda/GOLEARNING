package v1

import (
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golearning/inventory/internal/model"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

func (s *APISuite) TestListSuccess() {
	var (
		uuid    = gofakeit.UUID()
		name    = gofakeit.Word()
		country = gofakeit.Country()
		tag     = gofakeit.Word()

		req = &inventoryv1.ListPartsRequest{
			Filter: &inventoryv1.ListPartsFilter{
				Uuids:                 []string{uuid},
				Names:                 []string{name},
				Categorys:             []inventoryv1.Category{inventoryv1.Category_CATEGORY_ENGINE},
				ManufacturerCountries: []string{country},
				Tags:                  []string{tag},
			},
		}

		expectedFilter = model.ListPartsFilter{
			UUIDs:                 []string{uuid},
			Names:                 []string{name},
			Categories:            []model.Category{model.CategoryEngine},
			ManufacturerCountries: []string{country},
			Tags:                  []string{tag},
		}

		expectedParts = []*model.Part{
			{
				UUID: uuid,
				Name: name,
			},
		}
	)

	s.inventoryService.
		On("ListParts", s.ctx, expectedFilter).
		Return(expectedParts, nil)

	res, err := s.api.ListParts(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Len(res.GetParts(), 1)
	s.Require().Equal(uuid, res.GetParts()[0].GetUuid())
	s.Require().Equal(name, res.GetParts()[0].GetName())
}

func (s *APISuite) TestListServiceError() {
	var (
		serviceErr = gofakeit.Error()

		req = &inventoryv1.ListPartsRequest{
			Filter: &inventoryv1.ListPartsFilter{
				Uuids: []string{gofakeit.UUID()},
			},
		}

		expectedFilter = model.ListPartsFilter{
			UUIDs:      req.GetFilter().GetUuids(),
			Categories: []model.Category{},
		}
	)

	s.inventoryService.On("ListParts", s.ctx, expectedFilter).
		Return(([]*model.Part)(nil), serviceErr)

	res, err := s.api.ListParts(s.ctx, req)
	s.Require().Error(err)
	s.Require().Equal(codes.Internal, status.Code(err))
	s.Require().Nil(res)
}

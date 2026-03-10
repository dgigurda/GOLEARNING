package payment

import (
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appModel "golearning/payment/internal/model"
)

func (s *ServiceSuite) TestPaySuccess() {
	var (
		UserUUID              = gofakeit.UUID()
		orderUUID             = gofakeit.UUID()
		paymentMethod   int32 = 1
		transactionUUID       = gofakeit.UUID()
		req                   = appModel.PayOrderRequest{
			UserUUID:      UserUUID,
			OrderUUID:     orderUUID,
			PaymentMethod: paymentMethod,
		}

		expectedSavedPayment = appModel.Payment{
			OrderUUID:       orderUUID,
			UserUUID:        UserUUID,
			PaymentMethod:   paymentMethod,
			TransactionUUID: transactionUUID,
		}
	)
	s.service = &Service{
		repository: s.repo,
		newUUID:    func() string { return transactionUUID },
	}

	s.repo.
		On("Save", s.ctx, expectedSavedPayment).
		Return(nil)

	res, err := s.service.Pay(s.ctx, req)
	s.Require().NoError(err)
	s.Require().Equal(transactionUUID, res.TransactionUUID)
}

func (s *ServiceSuite) TestPayServiceError() {
	var (
		serviceErr            = gofakeit.Error()
		UserUUID              = gofakeit.UUID()
		orderUUID             = gofakeit.UUID()
		transactionUUID       = gofakeit.UUID()
		paymentMethod   int32 = 1
		req                   = appModel.PayOrderRequest{
			UserUUID:      UserUUID,
			OrderUUID:     orderUUID,
			PaymentMethod: paymentMethod,
		}

		expectedSavedPayment = appModel.Payment{
			OrderUUID:       orderUUID,
			UserUUID:        UserUUID,
			PaymentMethod:   paymentMethod,
			TransactionUUID: transactionUUID,
		}
	)

	s.service = &Service{
		repository: s.repo,
		newUUID:    func() string { return transactionUUID },
	}

	s.repo.
		On("Save", s.ctx, expectedSavedPayment).
		Return(serviceErr)

	res, err := s.service.Pay(s.ctx, req)
	s.Require().Error(err)
	s.Require().Equal(codes.Unknown, status.Code(err))
	s.Require().ErrorIs(err, serviceErr)
	s.Require().Equal(appModel.PayOrderResponse{}, res)

}

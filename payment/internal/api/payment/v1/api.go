package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golearning/payment/internal/model"
	"golearning/payment/internal/service"
	paymentv1 "golearning/shared/pkg/proto/payment/v1"
)

type API struct {
	paymentv1.UnimplementedPaymentServiceServer

	service service.PaymentService
}

func NewAPI(service service.PaymentService) *API {
	return &API{service: service}
}

func grpcError(err error) error {
	switch {
	case errors.Is(err, model.ErrInvalidOrderUUID), errors.Is(err, model.ErrInvalidUserUUID), errors.Is(err, model.ErrInvalidPaymentMethod):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func (a *API) PayOrder(ctx context.Context, request *paymentv1.PayOrderRequest) (*paymentv1.PayOrderResponse, error) {
	response, err := a.service.Pay(ctx, model.PayOrderRequest{
		OrderUUID:     request.GetOrderUuid(),
		UserUUID:      request.GetUserUuid(),
		PaymentMethod: int32(request.GetPaymentMethod()),
	})
	if err != nil {
		return nil, grpcError(err)
	}

	return &paymentv1.PayOrderResponse{TransactionUuid: response.TransactionUUID}, nil
}

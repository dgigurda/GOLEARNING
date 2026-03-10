package v1

import (
	"context"
	"errors"
	"net/http"

	"golearning/order/internal/model"
	"golearning/order/internal/service"
	orderv1 "golearning/shared/pkg/openapi/order/v1"
)

type API struct {
	service service.OrderService
}

func NewAPI(service service.OrderService) *API {
	return &API{service: service}
}

func (a *API) NewError(_ context.Context, err error) *orderv1.GenericErrorStatusCode {
	return &orderv1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: orderv1.GenericError{
			Code:    orderv1.NewOptInt(http.StatusInternalServerError),
			Message: orderv1.NewOptString(err.Error()),
		},
	}
}

func mapError(err error) (int, string) {
	switch {
	case errors.Is(err, model.ErrOrderNotFound):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, model.ErrOrderConflict):
		return http.StatusConflict, err.Error()
	case errors.Is(err, model.ErrInvalidCreateOrderRequest), errors.Is(err, model.ErrInvalidPayOrderRequest), errors.Is(err, model.ErrPartsNotFound):
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, err.Error()
	}
}

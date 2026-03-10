package v1

import (
	"context"
	"net/http"

	"golearning/order/internal/converter"
	orderv1 "golearning/shared/pkg/openapi/order/v1"
)

func (a *API) CreateOrder(ctx context.Context, request *orderv1.CreateOrderRequest) (orderv1.CreateOrderRes, error) {
	response, err := a.service.Create(ctx, converter.CreateRequestFromAPI(request))
	if err != nil {
		statusCode, message := mapError(err)
		if statusCode == http.StatusBadRequest {
			return &orderv1.BadRequestError{Code: statusCode, Message: message}, nil
		}
		return nil, err
	}

	return converter.CreateResponseToAPI(response), nil
}

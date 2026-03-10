package v1

import (
	"context"
	"net/http"

	"golearning/order/internal/converter"
	orderv1 "golearning/shared/pkg/openapi/order/v1"
)

func (a *API) PayOrder(ctx context.Context, request *orderv1.PayOrderRequest, params orderv1.PayOrderParams) (orderv1.PayOrderRes, error) {
	response, err := a.service.Pay(ctx, converter.PayRequestFromAPI(params.OrderUUID, request))
	if err != nil {
		statusCode, message := mapError(err)
		switch statusCode {
		case http.StatusNotFound:
			return &orderv1.NotFoundError{Code: statusCode, Message: message}, nil
		default:
			// PayOrder contract allows only NotFoundError for error responses.
			return &orderv1.NotFoundError{Code: http.StatusNotFound, Message: message}, nil
		}
	}

	return converter.PayResponseToAPI(response), nil
}

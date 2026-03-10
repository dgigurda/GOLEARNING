package v1

import (
	"context"
	"net/http"

	"golearning/order/internal/converter"
	orderv1 "golearning/shared/pkg/openapi/order/v1"
)

func (a *API) GetOrderByUuid(ctx context.Context, params orderv1.GetOrderByUuidParams) (orderv1.GetOrderByUuidRes, error) {
	order, err := a.service.Get(ctx, params.OrderUUID)
	if err != nil {
		statusCode, message := mapError(err)
		if statusCode == http.StatusNotFound {
			return &orderv1.NotFoundError{Code: statusCode, Message: message}, nil
		}
		return nil, err
	}

	return converter.OrderToAPI(order), nil
}

package v1

import (
	"context"
	"net/http"

	orderv1 "golearning/shared/pkg/openapi/order/v1"
)

func (a *API) CancelOrder(ctx context.Context, params orderv1.CancelOrderParams) (orderv1.CancelOrderRes, error) {
	if err := a.service.Cancel(ctx, params.OrderUUID); err != nil {
		statusCode, message := mapError(err)
		switch statusCode {
		case http.StatusNotFound:
			return &orderv1.NotFoundError{Code: statusCode, Message: message}, nil
		case http.StatusConflict:
			return &orderv1.ConflictError{Code: statusCode, Message: message}, nil
		default:
			return nil, err
		}
	}

	return &orderv1.CancelOrderNoContent{}, nil
}

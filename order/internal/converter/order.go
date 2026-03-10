package converter

import (
	orderv1 "golearning/shared/pkg/openapi/order/v1"

	"golearning/order/internal/model"
)

func CreateRequestFromAPI(request *orderv1.CreateOrderRequest) model.CreateOrderRequest {
	if request == nil {
		return model.CreateOrderRequest{}
	}
	return model.CreateOrderRequest{UserUUID: request.UserUUID, PartUUIDs: append([]string(nil), request.PartUuids...)}
}

func PayRequestFromAPI(orderUUID string, request *orderv1.PayOrderRequest) model.PayOrderRequest {
	if request == nil {
		return model.PayOrderRequest{OrderUUID: orderUUID}
	}
	return model.PayOrderRequest{OrderUUID: orderUUID, PaymentMethod: request.PaymentMethod}
}

func CreateResponseToAPI(response model.CreateOrderResponse) *orderv1.CreateOrderResponse {
	return &orderv1.CreateOrderResponse{OrderUUID: response.OrderUUID, TotalPrice: response.TotalPrice}
}

func OrderToAPI(order model.Order) *orderv1.Order {
	result := &orderv1.Order{
		OrderUUID:  order.OrderUUID,
		UserUUID:   order.UserUUID,
		PartUuids:  append([]string(nil), order.PartUUIDs...),
		TotalPrice: order.TotalPrice,
		Status:     string(order.Status),
	}
	if order.TransactionUUID != "" {
		result.TransactionUUID = orderv1.NewOptString(order.TransactionUUID)
	}
	if order.PaymentMethod != "" {
		result.PaymentMethod = orderv1.NewOptString(order.PaymentMethod)
	}
	return result
}

func PayResponseToAPI(response model.PayOrderResponse) *orderv1.PayOrderResponse {
	return &orderv1.PayOrderResponse{TransactionUUID: response.TransactionUUID}
}

package converter

import (
	"golearning/order/internal/model"
	repositoryModel "golearning/order/internal/repository/model"
)

func ToRepositoryOrder(order model.Order) repositoryModel.Order {
	return repositoryModel.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUUIDs:       append([]string(nil), order.PartUUIDs...),
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   order.PaymentMethod,
		Status:          repositoryModel.OrderStatus(order.Status),
	}
}

func ToModelOrder(order repositoryModel.Order) model.Order {
	return model.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUUIDs:       append([]string(nil), order.PartUUIDs...),
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   order.PaymentMethod,
		Status:          model.OrderStatus(order.Status),
	}
}

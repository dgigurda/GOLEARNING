package grpc

import (
	"google.golang.org/grpc"

	inventoryclient "golearning/order/internal/client/grpc/inventory/v1"
	paymentclient "golearning/order/internal/client/grpc/payment/v1"
)

func NewInventoryClient(conn grpc.ClientConnInterface) *inventoryclient.Client {
	return inventoryclient.NewClient(conn)
}

func NewPaymentClient(conn grpc.ClientConnInterface) *paymentclient.Client {
	return paymentclient.NewClient(conn)
}

package v1

import (
	"context"
	"strings"

	"google.golang.org/grpc"

	paymentv1 "golearning/shared/pkg/proto/payment/v1"
)

type Client struct {
	client paymentv1.PaymentServiceClient
}

func NewClient(conn grpc.ClientConnInterface) *Client {
	return &Client{client: paymentv1.NewPaymentServiceClient(conn)}
}

func NewFromGenerated(client paymentv1.PaymentServiceClient) *Client {
	return &Client{client: client}
}

func (c *Client) PayOrder(ctx context.Context, orderUUID, userUUID, paymentMethod string) (string, error) {
	method := paymentv1.PaymentMethod(paymentv1.PaymentMethod_value["PAYMENT_METHOD_"+strings.ToUpper(paymentMethod)])
	response, err := c.client.PayOrder(ctx, &paymentv1.PayOrderRequest{
		OrderUuid:     orderUUID,
		UserUuid:      userUUID,
		PaymentMethod: method,
	})
	if err != nil {
		return "", err
	}

	return response.GetTransactionUuid(), nil
}

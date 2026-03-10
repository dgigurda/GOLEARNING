package model

type Payment struct {
	OrderUUID       string
	UserUUID        string
	PaymentMethod   int32
	TransactionUUID string
}

type PayOrderRequest struct {
	OrderUUID     string
	UserUUID      string
	PaymentMethod int32
}

type PayOrderResponse struct {
	TransactionUUID string
}

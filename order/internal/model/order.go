package model

type OrderStatus string

const (
	OrderStatusPendingPayment OrderStatus = "PENDING_PAYMENT"
	OrderStatusPaid           OrderStatus = "PAID"
	OrderStatusCancelled      OrderStatus = "CANCELLED"
)

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartUUIDs       []string
	TotalPrice      float32
	TransactionUUID string
	PaymentMethod   string
	Status          OrderStatus
}

type CreateOrderRequest struct {
	UserUUID  string
	PartUUIDs []string
}

type CreateOrderResponse struct {
	OrderUUID  string
	TotalPrice float32
}

type PayOrderRequest struct {
	OrderUUID      string
	PaymentMethod  string
	UserUUID       string
	TransactionRef string
}

type PayOrderResponse struct {
	TransactionUUID string
}

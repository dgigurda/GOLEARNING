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

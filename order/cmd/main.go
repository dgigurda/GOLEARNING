package main

import (
	"context"
	// "errors"
	// "log"
	// "net"
	// "net/http"
	//"os"
	//"os/signal"
	"sync"
	//"syscall"
	"time"

	//"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	//"golang.org/x/mod/sumdb/storage"

	orderV1 "golearning/shared/pkg/openapi/order/v1"
)

const (
	httpPort = "8080"

	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type OrderStorage struct {
	mu     sync.RWMutex
	orders map[string]*orderV1.Order
}

func NewOrderService() *OrderStorage {
	return &OrderStorage{
		orders: make(map[string]*orderV1.Order),
	}
}

func (s *OrderStorage) GetOrder(uuid string) *orderV1.Order {
	s.mu.RLock()
	defer s.mu.RUnlock()

	order, ok := s.orders[uuid]
	if !ok {
		return nil
	}

	return order
}

type OrderHandler struct {
	storage *OrderStorage
}

func NewOrderHandler(storage *OrderStorage) *OrderHandler {
	return &OrderHandler{
		storage: storage,
	}
}

// CancelOrder invokes CancelOrder operation.
//
// Cancel order.
//
// POST /api/v1/orders/{order_uuid}/cancel
func (s *OrderHandler) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	return nil, nil
}

// CreateOrder invokes CreateOrder operation.
//
// Create order.
//
// POST /api/v1/orders/
func (s *OrderHandler) CreateOrder(_ context.Context, request *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {
	return nil, nil
}

// GetOrderByUuid invokes GetOrderByUuid operation.
//
// Get order data for a order uuid.
//
// GET /api/v1/order/{order_uuid}
func (s *OrderHandler) GetOrderByUuid(ctx context.Context, params orderV1.GetOrderByUuidParams) (orderV1.GetOrderByUuidRes, error) {
	return nil, nil
}

// PayOrder invokes PayOrder operation.
//
// Pay order.
//
// POST /api/v1/orders/{order_uuid}/pay
func (s *OrderHandler) PayOrder(ctx context.Context, request *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {
	return nil, nil
}

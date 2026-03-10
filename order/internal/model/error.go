package model

import "errors"

var (
	ErrInvalidCreateOrderRequest = errors.New("invalid create order request")
	ErrInvalidPayOrderRequest    = errors.New("invalid pay order request")
	ErrOrderNotFound             = errors.New("order not found")
	ErrOrderConflict             = errors.New("order state conflict")
	ErrExternalInventory         = errors.New("inventory service error")
	ErrExternalPayment           = errors.New("payment service error")
	ErrPartsNotFound             = errors.New("some parts not found")
)

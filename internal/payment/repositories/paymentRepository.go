package repositories

import (
	"errors"
	"tech-challenge-payment/internal/payment/types"
	"time"
)

var ErrPaymentNotFound = errors.New("Payment not found")

type PaymentRepository interface {
	CreatePayment(paymentID string, paymentData types.PaymentData) error
	GetPayment(paymentID string) (types.Payment, error)
	GetPaymentByOrderID(orderID string, timeThreshold time.Duration) (string, error)
}
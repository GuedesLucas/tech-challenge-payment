package repositories

import (
	"errors"
	"tech-challenge-payment/internal/payment/types"
	"time"
)

var ErrPaymentNotFound = errors.New("Payment not found")

type PaymentRepository interface {
	CreatePayment(paymentID string, paymentData types.PaymentData) error
	GetPaymentByPaymentID(paymentID string) (types.PaymentData, error)
	GetPaymentByOrderID(orderID int64, timeThreshold time.Duration) (string, error)
	SavePaymentStatus(paymentID string, status string) error
}

package payment

import "tech-challenge-payment/internal/payment/types"

type Service interface {
	CreatePayment(paymentData types.PaymentData) error
	GetPayment(paymentID string) (types.Payment, error)
}

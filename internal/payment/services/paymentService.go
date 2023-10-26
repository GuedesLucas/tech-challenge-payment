package payment

import (
	"database/sql"
	paymentRepository "tech-challenge-payment/internal/payment/repositories"
	"tech-challenge-payment/internal/payment/types"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GeneratePayment(paymentData types.PaymentData) (string, error)
	GetPayment(paymentID string) (types.Payment, error)
}

type paymentService struct {
	repository paymentRepository.PaymentRepository
}

func NewPaymentService(db *sql.DB) Service {
	repo := paymentRepository.NewPaymentRepository(db)
	return &paymentService{
		repository: repo,
	}
}

func (s *paymentService) GeneratePayment(paymentData types.PaymentData) (string, error) {

	existingPaymentID, err := s.repository.GetPaymentByOrderID(paymentData.OrderID, 5*time.Minute)
	if err != nil {
		return "", err
	}

	if existingPaymentID != "" {
		return existingPaymentID, nil
	}

	paymentID := uuid.New().String()
	if err := s.repository.CreatePayment(paymentID, paymentData); err != nil {
		return "", err
	}

	return paymentID, nil
}

// GetPayment implements Service.
func (*paymentService) GetPayment(paymentID string) (types.Payment, error) {
	panic("unimplemented")
}

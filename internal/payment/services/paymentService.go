package payment

import (
	"database/sql"
	"fmt"
	paymentRepository "tech-challenge-payment/internal/payment/repositories"
	"tech-challenge-payment/internal/payment/types"
	"tech-challenge-payment/internal/utils"
	api "tech-challenge-payment/pkg/api"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GeneratePayment(paymentData types.PaymentData) (string, error)
	GetPayment(paymentID string) (types.PaymentData, error)
	Payment(payData types.Payment) (string, error)
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
	paymentData.Status = "waiting"

	paymentID := uuid.New().String()
	if err := s.repository.CreatePayment(paymentID, paymentData); err != nil {
		return "", err
	}

	return paymentID, nil
}

// GetPayment implements Service.
func (s *paymentService) GetPayment(paymentID string) (types.PaymentData, error) {
	panic("unimplemented")
}

func (s *paymentService) Payment(payData types.Payment) (string, error) {
	existingPaymentID, err := s.repository.GetPayment(payData.ID)
	if err != nil {
		return "ok", err
	}
	r := utils.Randon()
	if r != "internal_failure" {
		if err := s.repository.SavePaymentStatus(existingPaymentID.ID, r); err != nil {
			WebHook(payData.ID)
			return "ok", err
		}
	}

	return "ok", err
}

func WebHook(paymentID string) string {
	apiClient := api.NewAPIClient("http://localhost:8080")

	_, err := apiClient.MakeRequest("GET", "/webhook/"+paymentID, nil, nil, nil)
	if err != nil {
		fmt.Printf("Erro ao fazer a chamada GET: %v\n", err)
	}
	return "ok"
}

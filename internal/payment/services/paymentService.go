package payment

import (
	"fmt"
	"tech-challenge-payment/config"
	paymentRepository "tech-challenge-payment/internal/payment/repositories"
	"tech-challenge-payment/internal/payment/types"
	"tech-challenge-payment/internal/utils"
	api "tech-challenge-payment/pkg/api"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	GeneratePayment(paymentData types.PaymentData) (string, error)
	GetPayment(paymentID string) (types.PaymentData, error)
	Payment(payData types.Payment) (string, error)
}

type paymentService struct {
	repository paymentRepository.PaymentRepository
	appConfig  config.AppConfig
}

func NewPaymentService(db *gorm.DB, appConfig config.AppConfig) Service {
	repo := paymentRepository.NewPaymentRepository(db)
	return &paymentService{
		repository: repo,
		appConfig:  appConfig,
	}
}

func (s *paymentService) GeneratePayment(paymentData types.PaymentData) (string, error) {

	existingPaymentID, _ := s.repository.GetPaymentByOrderID(paymentData.OrderID, 5*time.Minute)

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
	existingPaymentID, err := s.repository.GetPaymentByPaymentID(paymentID)
	if err != nil {
		return types.PaymentData{}, err
	}

	return existingPaymentID, nil
}

func (s *paymentService) Payment(payData types.Payment) (string, error) {
	existingPaymentID, err := s.repository.GetPaymentByPaymentID(payData.ID)
	if err != nil {
		return "internal_failure", err
	}
	r := utils.Randon()
	if r == "internal_failure" {
		return "internal_failure", err
	}
	if err := s.repository.SavePaymentStatus(existingPaymentID.ID, r); err != nil {
		return "internal_failure", err
	}
	s.WebHook(existingPaymentID.ID)

	return r, nil
}

func (s *paymentService) WebHook(paymentID string) string {
	apiClient := api.NewAPIClient(s.appConfig.Webhook.BaseURL)
	path := s.appConfig.Webhook.Path + paymentID

	_, err := apiClient.MakeRequest("GET", path, nil, nil, nil)
	if err != nil {
		fmt.Printf("Erro ao fazer a chamada GET: %v\n", err)
	}
	return "ok"
}

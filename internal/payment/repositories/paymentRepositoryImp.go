package repositories

import (
	"tech-challenge-payment/internal/payment/models"
	"tech-challenge-payment/internal/payment/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) CreatePayment(paymentID string, paymentData types.PaymentData) error {
	parsedUUID, _ := uuid.Parse(paymentID)
	payment := &models.Payment{
		ID:          parsedUUID,
		OrderID:     paymentData.OrderID,
		Amount:      paymentData.Amount,
		Status:      paymentData.Status,
		PaymentTime: time.Now(),
	}

	if err := r.db.Create(payment).Error; err != nil {
		return err
	}

	return nil
}

func (r *paymentRepository) GetPaymentByPaymentID(paymentID string) (types.PaymentData, error) {
	var payment models.Payment
	if err := r.db.Where("id = ?", paymentID).First(&payment).Error; err != nil {
		return types.PaymentData{}, err
	}

	return models.ToPaymentDTO(payment), nil
}

func (r *paymentRepository) GetPaymentByOrderID(orderID int64, timeThreshold time.Duration) (string, error) {
	currentTime := time.Now()
	var payment models.Payment
	if err := r.db.Where("order_id = ? AND payment_time >= ? AND status != 'fail'", orderID, currentTime.Add(-timeThreshold)).First(&payment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", ErrPaymentNotFound
		}
		return "", err
	}

	return payment.ID.String(), nil
}

func (r *paymentRepository) SavePaymentStatus(paymentID string, status string) error {
	if err := r.db.Model(&models.Payment{}).Where("id = ?", paymentID).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}

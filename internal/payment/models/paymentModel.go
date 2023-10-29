package models

import (
	"time"

	"tech-challenge-payment/internal/payment/types"

	"github.com/google/uuid"
)

type Payment struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	OrderID     int64
	Amount      float64
	PaymentTime time.Time
	Status      string
}

func ToPaymentDTO(model Payment) types.PaymentData {
	return types.PaymentData{
		ID:          model.ID.String(),
		OrderID:     model.OrderID,
		Amount:      model.Amount,
		PaymentTime: model.PaymentTime,
		Status:      model.Status,
	}
}

package types

import (
	"time"

	"github.com/go-playground/validator"
)

type PaymentData struct {
	ID          string    `json:"id"`
	OrderID     int64     `json:"order_id" validate:"required,gt=0"`
	Amount      float64   `json:"amount" validate:"required,gt=0"`
	PaymentTime time.Time `json:"payment_time"`
	Status      string    `json:"status"`
}

func NewPaymentData(orderID int64, amount float64) (*PaymentData, error) {
	paymentData := &PaymentData{
		OrderID: orderID,
		Amount:  amount,
	}

	// Use a biblioteca go-validator para validar os campos.
	validate := validator.New()
	if err := validate.Struct(paymentData); err != nil {
		return nil, err
	}

	return paymentData, nil
}

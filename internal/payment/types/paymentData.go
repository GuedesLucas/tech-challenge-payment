package types

import "time"

type PaymentData struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id"`
	Amount      float64   `json:"amount"`
	PaymentTime time.Time `json:"payment_time"`
}

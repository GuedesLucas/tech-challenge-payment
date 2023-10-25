package types

import "time"

type Payment struct {
	ID          string      `json:"id"`
	PaymentData PaymentData `json:"payment_data"`
	Status      string      `json:"status"`
	PaymentTime time.Time   `json:"payment_time"`
}

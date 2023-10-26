package repositories

import (
	"database/sql"
	"tech-challenge-payment/internal/payment/types"
	"time"
)

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) CreatePayment(paymentID string, paymentData types.PaymentData) error {
	_, err := r.db.Exec("INSERT INTO payments (payment_id, order_id, amount, payment_time) VALUES ($1, $2, $3, NOW())", paymentID, paymentData.OrderID, paymentData.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *paymentRepository) GetPayment(paymentID string) (types.Payment, error) {
	var payment types.Payment
	err := r.db.QueryRow("SELECT payment_id, order_id, amount, payment_time FROM payments WHERE payment_id = $1", paymentID).
		Scan(&payment.ID, &payment.PaymentData.OrderID, &payment.PaymentData.Amount, &payment.PaymentTime)
	if err != nil {
		return types.Payment{}, err
	}

	return payment, nil
}

func (r *paymentRepository) GetPaymentByOrderID(orderID string, timeThreshold time.Duration) (string, error) {
	currentTime := time.Now()
	rows, err := r.db.Query("SELECT payment_id, payment_time FROM payments WHERE order_id = $1 AND payment_time >= $2", orderID, currentTime.Add(-timeThreshold))
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var paymentID string
		var paymentTime time.Time
		if err := rows.Scan(&paymentID, &paymentTime); err != nil {
			return "", err
		}

		return paymentID, nil
	}

	return "", ErrPaymentNotFound
}

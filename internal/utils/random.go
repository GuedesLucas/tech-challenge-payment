package utils

import "math/rand"

func Randon() string {
	randomNumber := rand.Intn(100)

	var paymentStatus string

	if randomNumber < 95 {
		// Pagamento bem-sucedido (95% de sucesso)
		paymentStatus = "success"
	} else if randomNumber < 98 {
		// Falha no processamento do banco (3% de falha)
		paymentStatus = "fail"
	} else {
		// Falha interna (2% de falha)
		paymentStatus = "internal_failure"
	}
	return paymentStatus
}

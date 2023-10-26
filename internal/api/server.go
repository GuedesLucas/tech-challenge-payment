package api

import (
	"log"
	"tech-challenge-payment/config"
	payment "tech-challenge-payment/internal/payment/services"
)

func InitServices() payment.Service {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal("Erro ao carregar configurações do aplicativo:", err)
	}

	db, err := config.InitDatabase(appConfig.Database)
	if err != nil {
		log.Fatal("Erro ao inicializar o banco de dados:", err)
	}

	paymentService := payment.NewPaymentService(db)

	return paymentService
}

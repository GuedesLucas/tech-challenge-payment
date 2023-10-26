package api

import (
	"fmt"
	"log"

	config "tech-challenge-payment/config"
	"tech-challenge-payment/internal/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func StartServer() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal("Erro ao carregar configurações do aplicativo:", err)
	}

	apiPort := appConfig.API.Port

	r := gin.Default()

	paymentService := InitServices()

	routes.InitPaymentRoutes(r, paymentService)
	routes.InitSwaggerRoutes(r)

	address := fmt.Sprintf(":%d", apiPort)
	r.Run(address)
}

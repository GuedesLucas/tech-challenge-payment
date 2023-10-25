package api

import (
	"tech-challenge-payment/internal/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StartServer() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbURL := viper.GetString("database.url")
	apiPort := viper.GetInt("api.port")

	r := gin.Default()

	routes.InitPaymentRoutes(r)
	routes.InitSwaggerRoutes(r)

	r.Run(":" + apiPort)
}

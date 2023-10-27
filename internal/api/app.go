package api

import (
	"fmt"
	"log"
	config "tech-challenge-payment/config"
	docs "tech-challenge-payment/docs"
	"tech-challenge-payment/internal/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

//	@title			Swagger Payment-Moch API
//	@version		1.0
//	@description	This API is for Moch Payment to use on Tech-Challenge FIAP.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:7575

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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

	apiPort := appConfig.Server.Port

	r := gin.Default()

	paymentService := InitServices()

	docs.SwaggerInfo.Title = "Swagger Payment-Moch API"
	docs.SwaggerInfo.Description = "This API is for Moch Payment to use on Tech-Challenge FIAP."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	routes.InitPaymentRoutes(r, paymentService)
	routes.InitSwaggerRoutes(r)

	address := fmt.Sprintf(":%d", apiPort)
	r.Run(address)
}

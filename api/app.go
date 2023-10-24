package api

import (
	"github.com/gin-gonic/gin"
	"tech-challenge-payment/api/routes"
)

func StartServer() {
	r := gin.Default()

	routes.InitPaymentRoutes(r)

	r.Run(":7575")
}
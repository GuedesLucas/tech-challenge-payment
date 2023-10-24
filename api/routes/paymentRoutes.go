package routes

import (
	"github.com/gin-gonic/gin"
	"tech-challenge-payment/api/handlers"
)

func InitPaymentRoutes(r *gin.Engine) {
	pagamento := r.Group("/payment")
	{
		pagamento.GET("/", handlers.GetPayment)
		pagamento.POST("/", handlers.CreatePayment)
	}
}
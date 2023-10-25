package routes

import (
	"tech-challenge-payment/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitPaymentRoutes(r *gin.Engine) {
	pagamento := r.Group("/payment")
	{
		pagamento.GET("/", handlers.GetPayment)
		pagamento.POST("/", handlers.CreatePayment)
	}
}

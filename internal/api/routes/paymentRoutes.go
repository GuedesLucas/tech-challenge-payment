package routes

import (
	"tech-challenge-payment/internal/payment/handlers"
	payment "tech-challenge-payment/internal/payment/services"

	"github.com/gin-gonic/gin"
)

func InitPaymentRoutes(r *gin.Engine, paymentService payment.Service) {
	pagamento := r.Group("/payment")
	{
		pagamento.POST("/qrcode", func(c *gin.Context) {
			handlers.GeneratePaymentCode(c, paymentService)
		})
	}
}

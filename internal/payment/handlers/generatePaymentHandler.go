package handlers

import (
	"net/http"
	payment "tech-challenge-payment/internal/payment/services"
	"tech-challenge-payment/internal/payment/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

func GeneratePaymentCode(c *gin.Context, paymentService payment.Service) {
	var request types.PaymentData
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentID, err := paymentService.GeneratePayment(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to generate the payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment_id": paymentID})
}

func Payment(c *gin.Context, paymentService payment.Service) {
	var request types.Payment
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentID, err := paymentService.Payment(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to generate the payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": paymentID})
}

func GetPayment(c *gin.Context, paymentID string, paymentService payment.Service) {
	_, err := uuid.Parse(paymentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de pagamento inválido"})
		return
	}

	payment, err := paymentService.GetPayment(paymentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pagamento não encontrado"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

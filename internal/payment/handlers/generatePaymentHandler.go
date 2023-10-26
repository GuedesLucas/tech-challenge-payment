package handlers

import (
	"net/http"
	payment "tech-challenge-payment/internal/payment/services"
	"tech-challenge-payment/internal/payment/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

// @Summary Generate a payment code.
// @Description Generate a payment code based on the provided data.
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body PaymentData true "Data for generating the payment code"
// @Success 200 {object} PaymentCodeResponse "ID of the successfully generated payment"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 500 {object} ErrorResponse "Internal error generating the payment code"
// @Router /payment/generate [post]
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

// @Summary Process a payment.
// @Description Process a payment using the provided payment data.
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body Payment true "Payment data for processing"
// @Success 200 {object} PaymentStatusResponse "Payment status"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 500 {object} ErrorResponse "Internal error processing the payment"
// @Router /payment/process [post]
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

// @Summary Get payment details.
// @Description Get payment details by ID.
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} Payment "Payment details"
// @Failure 400 {object} ErrorResponse "Invalid payment ID"
// @Failure 404 {object} ErrorResponse "Payment not found"
// @Router /payment/{id} [get]
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

package handlers

import (
	"net/http"
	payment "tech-challenge-payment/internal/payment/services"
	"tech-challenge-payment/internal/payment/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// GeneratePaymentCode recebe uma solicitação HTTP para gerar um código de pagamento.
// @Summary Gera um código de pagamento.
// @Description Gera um código de pagamento com base nos dados fornecidos.
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body GeneratePaymentCodeRequest true "Dados para gerar o código de pagamento"
// @Success 200 {object} PaymentCodeResponse "ID do pagamento gerado com sucesso"
// @Failure 400 {object} ErrorResponse "Requisição inválida"
// @Failure 500 {object} ErrorResponse "Erro interno ao gerar o código de pagamento"
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

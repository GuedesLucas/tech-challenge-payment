package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPayment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Endpoint GET de pagamento",
	})
}

func CreatePayment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Endpoint POST de pagamento",
	})
}
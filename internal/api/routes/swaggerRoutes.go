package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/http-swagger"
	_ "tuaysa.com/docs"
)

func SwaggerRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))
}

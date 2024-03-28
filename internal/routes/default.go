package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tuaysa.com/pkg/config"
	"tuaysa.com/pkg/response"
)

// For /v0/
func DefaultRoutes(group *gin.RouterGroup) {
	group.GET("/", func(c *gin.Context) {
		response.Success(c, http.StatusOK, "Tuaysa", gin.H{
			"message": "Welcome to the Tuaysa REST API v0",
			"version": config.AppConfig().App.AppVersion,
			"author":  "Gowthaman Ravindrathas",
		})
	})
}

// For /
func NoVersionRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		response.Success(c, http.StatusOK, "Tuaysa", gin.H{
			"message": "Welcome to the Tuaysa REST API v0",
			"version": config.AppConfig().App.AppVersion,
			"author":  "Gowthaman Ravindrathas",
			"api":     "v0",
		})
	})
}

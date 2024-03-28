package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tuaysa.com/internal/routes"
	"tuaysa.com/pkg/config"
	"tuaysa.com/pkg/db"
)

// @title Tuaysa API
// @version 0
// @description This is the REST API for Tuaysa.
// @host api.tuaysa.com
// @BasePath /v0
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	db.ConnectToMongoDB()
	defer db.DisconnectFromMongoDB()

	router := gin.Default()

	// Routes
	routes.SwaggerRoutes(router)
	routes.NoVersionRoutes(router)

	// Version controlled routes
	versionControlled := router.Group("/" + config.AppConfig().App.ApiVersion)
	{
		routes.DefaultRoutes(versionControlled)
	}

	router.Run(config.AppConfig().App.Host + config.AppConfig().App.Port)
}

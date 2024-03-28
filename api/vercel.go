package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tuaysa.com/internal/routes"
	"tuaysa.com/pkg/config"
	"tuaysa.com/pkg/db"
)

// @title Tuaysa API
// @version 1
// @description This is the REST API for Tuaysa.
// @host api.tuaysa.com
// @BasePath /v0
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Handler(w http.ResponseWriter, r *http.Request) {
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

	router.ServeHTTP(w, r)
}

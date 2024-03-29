package routes

import (
	"github.com/gin-gonic/gin"
	"tuaysa.com/middleware"
	"tuaysa.com/services/product"
)

// GROUP: /product
func ProductRoutes(group *gin.RouterGroup) {
	// Setup the repository & handler
	productRepo := product.NewProductRepository()
	productHandler := product.NewProductHandler(productRepo)

	productGroup := group.Group("/product")

	// --- PUBLIC ROUTES ---
	productGroup.POST("/create", func(c *gin.Context) {
		productHandler.CreateProduct(c)
	})

	productGroup.POST("/createMany", func(c *gin.Context) {
		productHandler.CreateManyProduct(c)
	})

	productGroup.POST("/all", func(c *gin.Context) {

	})

	// --- PROTECTED ROUTES ---
	productGroup.Use(middleware.JWTAuthMiddleware())
	{

	}

}

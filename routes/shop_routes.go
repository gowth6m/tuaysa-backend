package routes

import (
	"github.com/gin-gonic/gin"
	"tuaysa.com/middleware"
	"tuaysa.com/services/shop"
)

// GROUP: /shop
func ShopRoutes(group *gin.RouterGroup) {
	// Setup the repository & handler
	shopRepo := shop.NewShopRepository()
	shopHandler := shop.NewShopHandler(shopRepo)

	shopGroup := group.Group("/shop")

	// --- PUBLIC ROUTES ---
	shopGroup.POST("/create", func(c *gin.Context) {
		shopHandler.CreateShop(c)
	})

	shopGroup.GET("/all", func(c *gin.Context) {
		shopHandler.GetAllShops(c)
	})

	// --- PROTECTED ROUTES ---
	shopGroup.Use(middleware.JWTAuthMiddleware())
	{

	}

}

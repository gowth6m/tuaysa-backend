package routes

import (
	"github.com/gin-gonic/gin"
	"tuaysa.com/middleware"
)

func UserRoutes(group *gin.RouterGroup) {
	group.GET("/user", func(c *gin.Context) {
		userGroup := group.Group("/user")

		// Public routes
		userGroup.POST("/create", func(c *gin.Context) {
			// user.HandleCreateUser(c)
		})

		userGroup.POST("/login", func(c *gin.Context) {
			// user.HandleLogin(c)
		})

		// Private routes
		userGroup.Use(middleware.JWTAuthMiddleware())
		{
			userGroup.GET("/current", func(c *gin.Context) {
				// user.HandleGetCurrentUser(c)
			})
			
			
		}

	})
}

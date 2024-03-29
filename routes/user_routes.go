package routes

import (
	"github.com/gin-gonic/gin"
	"tuaysa.com/middleware"
	"tuaysa.com/services/user"
)

// GROUP: /user
func UserRoutes(group *gin.RouterGroup) {
	// Setup the repository & handler
	userRepo := user.NewUserRepository()
	userHandler := user.NewUserHandler(userRepo)

	userGroup := group.Group("/user")

	// --- PUBLIC ROUTES ---
	userGroup.POST("/create", func(c *gin.Context) {
		userHandler.CreateUser(c)
	})

	userGroup.GET("/all", func(c *gin.Context) {
		userHandler.GetAllUsers(c)
	})

	// --- PROTECTED ROUTES ---
	userGroup.Use(middleware.JWTAuthMiddleware())
	{

	}

}

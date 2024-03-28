package response

import (
	"github.com/gin-gonic/gin"
)

// ResponseData defines the structure of the API response
type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSON creates a standardized JSON response
func JSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, ResponseData{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

// Error is a convenience function to send error messages
func Error(c *gin.Context, statusCode int, message string) {
	JSON(c, statusCode, message, nil)
}

// Success is a convenience function to send success responses
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	JSON(c, statusCode, message, data)
}

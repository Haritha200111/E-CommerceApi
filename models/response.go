package models

import (
	"ecommerce/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string
	Status  int
	Data    interface{}
	Error   *error.Error
}

func CreateErrorResponse(c *gin.Context, statusCode int, message string, err *error.Error) {
	c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
		Error:   err,
	})
}

// Function to create a success response
func CreateSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

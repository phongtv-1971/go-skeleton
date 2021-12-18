package util

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
		"success": false,
	}
}

func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"data": data,
		"success": true,
	}
}

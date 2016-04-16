package controller

import (
	"github.com/gin-gonic/gin"
)

func abortWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}

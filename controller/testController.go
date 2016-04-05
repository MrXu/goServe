package controller

import (
	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) {
	c.String(200, "from test controller/n")
}

func TestJsonController(c *gin.Context) {
	c.JSON(200, gin.H{
		"name":   "jason",
		"action": "work",
	})
}

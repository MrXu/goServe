package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goServe/auth"
)

func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"greeting":"hello go",	
	})
}

func TestJsonController(c *gin.Context) {
	userAccount, exist := c.Get(auth.User)
	if !exist {
		abortWithError(c, http.StatusBadRequest, "user not found")
	}
	user := userAccount.(*auth.UserAccount)
	c.JSON(200, gin.H{
		"greeting":  "hello world",
		"email":user.Id,
	})
}

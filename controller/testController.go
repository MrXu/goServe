package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gorgeousServer/auth"
)

func TestController(c *gin.Context) {
	c.String(200, "from test controller/n")
}

func TestJsonController(c *gin.Context) {
	userAccount, exist := c.Get(auth.User)
	if !exist {
		abortWithError(c, http.StatusBadRequest, "user not found")
	}
	user := userAccount.(*auth.UserAccount)
	c.JSON(200, gin.H{
		"name":   "jason",
		"action": "work",
		"email":user.Id,
	})
}

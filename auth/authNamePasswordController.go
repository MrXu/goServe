package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginCredential struct{
	Email		string	`form:"email" json:"email" binding:"required"`
	Password	string	`form:"password" json:"password" binding:"required"`
}


func LoginUserWithPassword(c *gin.Context){
	var loginJson LoginCredential
	if c.BindJSON(loginJson) == nil {
		if loginJson.Email == "xw" && loginJson.Password =="xw"{
			c.JSON(http.StatusOK, gin.H{"authenticated":"true"})
		}else{
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated":"false"})
		}
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{"authenticated":"false"})
	}
}
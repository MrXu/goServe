package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gorgeousServer/mongodb"
	"gopkg.in/mgo.v2"
	"time"
)

type LoginCredential struct{
	Email		string	`form:"email" json:"email" binding:"required"`
	Password	string	`form:"password" json:"password" binding:"required"`
}

type SignUpCredential struct{
	Email  		string 	`form:"email" json:"email" binding:"required"`
	Password	string	`form:"password" json:"password" binding:"required"`
}

func LoginUserWithEmail(c *gin.Context){
	var loginJson LoginCredential

	if c.BindJSON(&loginJson) == nil {

		var user *UserAccount
		user, err := getUserByEmail(loginJson.Email, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated":"false"})
		}

		passwordValidErr := safeComparePassword(user.Password, []byte(loginJson.Password))
		if passwordValidErr != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated":"false"})	
		}

		tokenString, tokenErr := GenerateToken(loginJson.Email)

		if tokenErr != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated":"false"})		
		}

		c.JSON(http.StatusOK, gin.H{"authenticated":"true","token":tokenString})

		// if loginJson.Email == "xw" && loginJson.Password =="xw"{
		// 	c.JSON(http.StatusOK, gin.H{"authenticated":"true"})
		// }else{
		// 	c.JSON(http.StatusUnauthorized, gin.H{"authenticated":"false"})
		// }
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"unauthorized"})
	}
}


func SignUpWithEmail(c *gin.Context) {
	var signUpJson SignUpCredential	
	if c.BindJSON(&signUpJson) == nil {
		
		if validateEmail(signUpJson.Email, c) && validatePassword(signUpJson.Password){
			hash, hasherr := hashPassword(signUpJson.Password)
			if hasherr != nil{
				c.JSON(http.StatusBadRequest, gin.H{"error":"sign up fail"})
			}
			db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
			err := db.C(CollectionUserAccount).Insert(&UserAccount{
				Id:signUpJson.Email,
				Password:hash,
				CreatedOn:int64(time.Now().Second()),
				UpdatedOn:int64(time.Now().Second())})
			if err != nil{
				c.JSON(http.StatusBadRequest, gin.H{"error":"sign up fail"})			
			}

			// sent email

			// login and sent back token

			c.JSON(http.StatusOK, gin.H{"":""})
		}else{
			c.JSON(http.StatusBadRequest, gin.H{"error":"sign up fail"})	
		}

	}else{
		c.JSON(http.StatusBadRequest, gin.H{"error":"sign up fail"})
	}
}


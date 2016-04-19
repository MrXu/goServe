package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goServe/mongodb"
	"gopkg.in/mgo.v2"
)

const (
	EmailVerifyCode  	string = "vcode"
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
		// db access
		db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
		user, err := GetUserByEmail(loginJson.Email, db)

		if err != nil {
			abortWithError(c, http.StatusUnauthorized, "authentication fail")
			return
		}

		passwordValidErr := safeComparePassword(user.Password, []byte(loginJson.Password))
		if passwordValidErr != nil{
			abortWithError(c, http.StatusUnauthorized, "authentication fail")
			return
		}

		tokenString, tokenErr := GenerateToken(loginJson.Email)

		if tokenErr != nil{
			abortWithError(c, http.StatusUnauthorized, "authentication fail")
			return	
		}

		c.JSON(http.StatusOK, gin.H{"authenticated":"true","token":tokenString})

		
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"unauthorized"})
	}
}


func SignUpWithEmail(c *gin.Context) {
	var signUpJson SignUpCredential	
	if c.BindJSON(&signUpJson) == nil {
		
		if validateEmail(signUpJson.Email, c) && validatePassword(signUpJson.Password){
			// db access
			db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
			err := InsertUserAccount(signUpJson.Email, signUpJson.Password, db)

			if err != nil{
				abortWithError(c, http.StatusBadRequest, "signup fail")			
				return
			}

			// sent email
			sendRegistrationConfirmationEmail(signUpJson.Email,signUpJson.Email,c)

			// login and sent back token
			tokenString, tokenErr := GenerateToken(signUpJson.Email)
			if tokenErr != nil{
				abortWithError(c, http.StatusUnauthorized, "authentication fail")
				return	
			}
			c.JSON(http.StatusOK, gin.H{"status":"sign up success","authenticated":"true","token":tokenString})

		}else{
			c.JSON(http.StatusBadRequest, gin.H{"error":"sign up fail"})	
		}

	}else{
		c.JSON(http.StatusBadRequest, gin.H{"error":"sign up fail"})
	}
}


func ActivateAccountAfterEmailSignup(c *gin.Context) {
	vcode := c.Query(EmailVerifyCode)
	db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
	result,err := GetEmailConfirm(vcode,db)

	// err getting record
	if err != nil {
		abortWithError(c, http.StatusBadRequest, "fail to get code")
		return
	}

	if result.Used{
		abortWithError(c, http.StatusBadRequest, "code used ")	
		return
	}

	// vcode used or vcode expire
	if isTimeStampExpired(result.ExpireAt){
		abortWithError(c, http.StatusBadRequest, "code expired")	
		return
	}


	updateErr := RemoveEmailConfirm(vcode, db)
	if updateErr != nil {
		abortWithError(c, http.StatusBadRequest, "fail to update")	
		return
	}

	c.JSON(http.StatusOK, gin.H{"activated":"true"})
}


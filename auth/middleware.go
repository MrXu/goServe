package auth

import (
	"net/http"
	"goServe/mongodb"
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)


func TokenAuthMiddleware(c *gin.Context) {

	token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(JWTKEY))

			return b, nil
		})

	if err != nil {
		abortWithError(c, http.StatusUnauthorized, "Invaild User Token")
		return
	}


	if token.Claims[USERID] != nil{
		db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
		user, getUserErr := GetUserByEmail(token.Claims[USERID].(string),db)
		if getUserErr != nil{
			abortWithError(c, http.StatusUnauthorized, "Invalid User Token")
		}
		c.Set(User, user)
	}else{
		abortWithError(c, http.StatusUnauthorized, "Invalid User Token")	
	}

	c.Next()
}

func validateEmailLoginToken(token *jwt.Token) {
	
}


package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/dchest/uniuri"
	"time"
	"goServe/mongodb"
	"gopkg.in/mgo.v2"
)


func validateEmail(userId string, c *gin.Context) bool {
	db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
	_,err := GetUserByEmail(userId, db)
	if err != nil{	
		return true
	}else{
		return false
	}
}

// hash password with bcypt
func hashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}
	return hash,nil
}

func safeComparePassword(hash []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err
}

func validatePassword(password string) bool{
	if len(password) < 5{
		return false
	}
	return true
}

func generateRandomUri() string{
	b := uniuri.NewLen(64)
	return b
}


func abortWithError(c *gin.Context, code int, message string) {
	c.Header("WWW-Authenticate", "JWT realm="+Realm)
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}

func isTimeStampExpired(timestamp int64) bool{
	now := time.Now()
	expireAt := time.Unix(timestamp,0)
	return expireAt.Before(now)
}

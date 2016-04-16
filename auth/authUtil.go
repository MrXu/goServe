package auth

import (
	"github.com/gin-gonic/gin"
	"gorgeousServer/mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func getUserByEmail(userId string, c *gin.Context) (*UserAccount, error){
	db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
	result := &UserAccount{}
	err := db.C(CollectionUserAccount).Find(bson.M{"_id":userId}).One(&result)

	return result,err

}

func validateEmail(userId string, c *gin.Context) bool {
	db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)

	result := &UserAccount{}

	err := db.C(CollectionUserAccount).Find(bson.M{"_id":userId}).One(&result)
	if err != nil{	
		return true
	}else{
		return false
	}
}

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

func abortWithError(c *gin.Context, code int, message string) {
	c.Header("WWW-Authenticate", "JWT realm="+Realm)
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}
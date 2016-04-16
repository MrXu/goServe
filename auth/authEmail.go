package auth

import (
	"github.com/gin-gonic/gin"
	"goServe/mongodb"
	"gopkg.in/mgo.v2"
	"time"
)


func sendRegistrationConfirmationEmail(email string, userId string, c *gin.Context){
	context := c.Copy()
	randomToken := generateRandomToken()
	db := context.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
	err := db.C(CollectionEmailConfirmation).Insert(&emailConfirmation{
		UserId:userId,
		Token:randomToken,
		Used: false,
		ExpireAt:int64(time.Now().Second()),
		})
	if err!=nil {
		return
	}
}

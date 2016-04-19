package auth

import (
	"github.com/gin-gonic/gin"
	"goServe/mongodb"
	"gopkg.in/mgo.v2"
	"time"
	emailWorker "goServe/email"
)


func sendRegistrationConfirmationEmail(email string, userId string, c *gin.Context){
	context := c.Copy()
	randomToken := generateRandomUri()
	db := context.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
	err := db.C(CollectionEmailConfirmation).Insert(&EmailConfirmation{
		UserId:userId,
		Token:randomToken,
		Used: false,
		ExpireAt:int64(time.Now().Second()),
		})
	if err!=nil {
		return
	}

	go func() {
		emailWorker.SendAnEmail(email,randomToken)
	}()

}
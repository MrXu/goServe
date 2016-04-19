package auth

import (
	"github.com/gin-gonic/gin"
	"goServe/mongodb"
	"gopkg.in/mgo.v2"
	emailWorker "goServe/email"
)


func sendRegistrationConfirmationEmail(email string, userId string, c *gin.Context){
	db := c.MustGet(mongodb.DBMiddlewareName).(*mgo.Database)
	randomToken, err := InsertEmailConfirm(userId, db)
	if err!=nil {
		return
	}

	go func() {
		emailWorker.SendAnEmail(email,randomToken)
	}()

}

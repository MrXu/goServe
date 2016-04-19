package auth

import (
	
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionEmailConfirmation  	string 	=	"email_confirmations_tmp"
	CollectionPasswordChange 		string 	=	"password_change_tmp"
)

const (
	emailConfirmExpireDuration   		time.Duration 	= time.Hour*7
	passwordChangeExpireDuration 		time.Duration 	= time.Hour*1
)


type tempOpsToken struct{
	UserId 			string
	Token 			string			`json:"token" bson:"token"`
	Used 			bool
	ExpireAt 		int64
}


type EmailConfirmation struct{
	UserId 			string
	Token 			string			`json:"token" bson:"token"`
	Used 			bool			`json:"used" bson:"used"`
	ExpireAt 		int64
}

type PasswordChange struct{
	UserId 			string
	Token 			string			`json:"token" bson:"token"`
	Used 			bool			`json:"used" bson:"used"`
	ExpireAt 		int64
}

func InsertEmailConfirm(userId string, db *mgo.Database) (string,error) {
	randomToken := generateRandomUri()
	err := db.C(CollectionEmailConfirmation).Insert(&EmailConfirmation{
		UserId:userId,
		Token:randomToken,
		Used: false,
		ExpireAt:int64(time.Now().Add(emailConfirmExpireDuration).Unix()),
		})
	return randomToken, err
}

func GetEmailConfirm(vcode string, db *mgo.Database) (*EmailConfirmation,error) {
	result := &EmailConfirmation{}
	err := db.C(CollectionEmailConfirmation).Find(bson.M{"token":vcode}).One(&result)
	return result,err
}

func SetEmailConfirmUsed(vcode string, db *mgo.Database) error{
	err := db.C(CollectionEmailConfirmation).Update(bson.M{"token":vcode}, bson.M{"$set":bson.M{"used":true}})
	return err
}

func RemoveEmailConfirm(vcode string, db *mgo.Database) error {
	err := db.C(CollectionEmailConfirmation).Remove(bson.M{"token":vcode})
	return err
}






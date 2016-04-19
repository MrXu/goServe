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
	emailConfirmExpireTime   		time.Duration 	= time.Hour*7
	passwordChangeExpireTime 		time.Duration 	= time.Hour*1
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


func GetEmailConfirm(vcode string, db *mgo.Database) (*EmailConfirmation,error) {
	result := &EmailConfirmation{}
	err := db.C(CollectionEmailConfirmation).Find(bson.M{"token":vcode}).One(&result)
	return result,err
}

func SetEmailConfirmUsed(vcode string, db *mgo.Database) error{
	err := db.C(CollectionEmailConfirmation).Update(bson.M{"token":vcode}, bson.M{"$set":bson.M{"used":true}})
	return err
}







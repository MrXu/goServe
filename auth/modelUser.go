package auth

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	User 			string = "user"
)

const (
	CollectionUserAccount = "accounts"
)

const (
	FACEBOOK 		string = 	"facebook"
	GOOGLE 			string =	"google"
)

type UserAccount struct{
	Id 				string				`json:"id" bson:"_id"`
	Password 		[]byte		
	Profile 		UserProfile
	SocialAuth  	SocialAuth
	CreatedOn		int64				
	UpdatedOn		int64
	Active			bool 
}

type UserProfile struct{
	Email			string				`json:"email" bson:"email"`
	Name			string				`json:"name" bson:"name"`

}

type SocialAuth struct{
	service			string				`json:"service" bson:"service"`
	token			string
}



func GetUserByEmail(userId string, db *mgo.Database) (*UserAccount, error){
	result := &UserAccount{}
	err := db.C(CollectionUserAccount).Find(bson.M{"_id":userId}).One(&result)

	return result,err
}

func InsertUserAccount(email string, password string, db *mgo.Database) error {

	hash, hasherr := hashPassword(password)
	if hasherr != nil{
		return hasherr
	}

	err := db.C(CollectionUserAccount).Insert(&UserAccount{
		Id:			email,
		Password:	hash,	// hash
		CreatedOn:int64(time.Now().Second()),
		UpdatedOn:int64(time.Now().Second()),
		Active:false})

	return err

}



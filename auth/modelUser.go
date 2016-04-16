package auth

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

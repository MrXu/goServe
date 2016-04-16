package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JWTKEY  	string  		= "testingkey"
	ExpireTime 	time.Duration 	= time.Hour * 24 * 7
	Realm		string 			= "jwt auth"
	USERID		string 			= "userId"
	EXPIREAT 	string 			= "expireAt"
)

func GenerateToken(userId string) (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims[USERID] = userId
	token.Claims[EXPIREAT] = time.Now().Add(ExpireTime)
	tokenString, err := token.SignedString([]byte(JWTKEY))
	return tokenString,err
}

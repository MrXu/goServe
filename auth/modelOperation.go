package auth

import (
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
	Token 			[]byte
	Used 			bool
	ExpireAt 		int64
}


type emailConfirmation struct{
	tempOpsToken
}

type passwordChange struct{
	tempOpsToken
}

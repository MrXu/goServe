package email

import (
	"gopkg.in/gomail.v2"
	"goServe/config"
	"log"
)

var (
	Email   	config.Email
)

func ConfigEmail() {
	email,err := config.GetAnEmail()
	if err != nil {
		log.Fatal(err)
	}
	Email = email
}

func SendEmail(to string, msg string) {
	
	emailAuth := smtp.PlainAuth("", Email.Address, Email.Password, Email.Host)

	err := smtp.SendMail(Email.Host+":"+Email.Port, emailAuth, Email.Address, []string{to}, []byte(msg))

	if err != nil {
		log.Println(err)
	}

}

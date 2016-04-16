package email

import (
	"net/smtp"
	"goServe/config"
	"log"
)

var (
	Email   	*config.Email = config.GetAnEmail()
)

func SendEmail(to string, msg string) {
	emailAuth := smtp.PlainAuth("", Email.Address, Email.Password, Email.Host)
	err := smtp.SendMail(Email.Host, emailAuth, Email.Address, to, msg)

	if err != nil {
		log.Fatal(err)
	}

}

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
	err := smtp.SendMail(Email.Host+":"+Email.Port, emailAuth, Email.Address, []string{to}, []byte(msg))

	if err != nil {
		log.Fatal(err)
	}

}

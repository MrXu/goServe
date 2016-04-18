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


func SendAnEmail(to string, msg string){
	m := gomail.NewMessage()
	m.SetHeader("From", Email.Address)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Email</b>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(Email.Host, Email.Port, Email.Address, Email.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
	    log.Println(err)
	}
}
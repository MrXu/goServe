package config

import (
	"os"
	"encoding/json"
	"fmt"
	"errors"
)

var (
	Config 		Configuration
)

type Configuration struct{
	MongoDBUrl 	string
	Emails		[]Email
	Jwtkey		string

	Facebook 	FacebookClient
}

type Email struct{
	Address		string
	Password	string
	Host		string
	Port 		int
}

type FacebookClient struct{
	ClientID 		string
	ClientSecret 	string
}



func GetConfig(){
	file,err:=os.Open("./config_prod.json")
	if err != nil{
		panic("Configuration file missing!")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Configuration{}
	decodeErr := decoder.Decode(&config)
	if decodeErr!=nil {
		panic("Configuration file error")
	}
	Config = config
	fmt.Println("Got configuration")
	// remove the two line below
	fmt.Println("JWTKey: ",Config.Jwtkey)
	fmt.Println("Emails: ",Config.Emails)
	fmt.Println("Facebook:",Config.Facebook.ClientID)
}

func GetJwtKey() string{
	return Config.Jwtkey
}

func GetAnEmail() (Email,error) {
	var email Email
	if len(Config.Emails) < 1{
		return email,errors.New("email missing")
	}
	return Config.Emails[0], nil
}

func GetMongoDBUrl() string{
	return Config.MongoDBUrl
}

func GetFacebookClientId() string{
	return Config.Facebook.ClientID
}

func GetFacebookClientSecret() string {
	return Config.Facebook.ClientSecret
}
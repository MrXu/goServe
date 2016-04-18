package config

import (
	"os"
	"encoding/json"
	"fmt"
)

var (
	Config 		Configuration
)

type Configuration struct{
	MongoDBUrl 	string
	Emails		[]Email
	Jwtkey		string
}

type Email struct{
	Address		string
	Password	string
	Host		string
	Port 		string
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
}

func GetJwtKey() string{
	return Config.Jwtkey
}

func GetAnEmail() *Email {
	if len(Config.Emails) < 1{
		return nil
	}
	return &(Config.Emails[0])
}

func GetMongoDBUrl() string{
	return Config.MongoDBUrl
}
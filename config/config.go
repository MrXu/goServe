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
	Emails		[]Email
	Jwtkey		string
}

type Email struct{
	Address		string
	Password	string
	Host		string
}

func GetConfig(){
	file,err:=os.Open("./config.json")
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

func GetAnEmail() (*Email, error) {
	if len(Config.Emails) < 1{
		return nil, errors.New("No email configured")
	}
	return &(Config.Emails[0]),nil
}
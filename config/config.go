package config

import (
	"os"
	"encoding/json"
)

type Configuration struct{
	emails		[]email
	jwtkey		string
}

type email struct{
	address		string
	password	string
}

func GetConfig() Configuration{
	file,err:=os.Open("config.json")
	if err != nil{
		panic("Configuration file missing!")
	}
	decoder := json.NewDecoder(file)
	config := Configuration{}
	decodeErr := decoder.Decode(config)
	if decoder!=nil {
		panic("Configuration file error")
	}
	return config
}
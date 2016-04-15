package mongodb

import (
	"os"
	"fmt"
	"gopkg.in/mgo.v2"
)

var (
	//mongo session
	Session *mgo.Session

	//mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// default mongodb url
	MongoDBUrl = "mongodb://localhost:27017/"
)

func Connect() {
	uri := os.Getenv("MONGODB_URL")
	if len(uri)==0 {
		uri = MongoDBUrl
	}

	// parse uri into mgo.DialInfo
	mongo, err := mgo.ParseURL(uri)

	s,err := mgo.Dial(uri)
	if err != nil{
		fmt.Printf("Cannot connect to mongodb, go error %v\n", err)
		panic(err.Error())
	}

	// set safe mode of mongodb
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to ",uri)
	Session = s
	Mongo = mongo
}

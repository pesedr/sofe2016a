package repo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	testDatabase = "test"
	mongoPort    = "mongodb://localhost"
)

func getSession() *mgo.Session {
	log.Println("Starting mongo session ", "Address: ", mongoPort)
	session, err := mgo.Dial(mongoPort)
	if err != nil {
		log.Fatalln("Couldn't connect to Database. ", err.Error())
	}

	return session
}

func stringIDtoObjectID(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(id) {
		log.Println("ID is not mongo objectID, ", "ID: ", id)
		return "", nil
	}
	return bson.ObjectIdHex(id), nil

}

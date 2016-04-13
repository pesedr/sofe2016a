package repo

import (
	"fmt"
	"log"

	"github.com/pesedr/sofe2016a/errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	testDatabase = "test"
	mongoPort    = "mongodb://localhost"
)

func init() {
	session := getSession()

	// Any new Repository interfaces that need the db should be initialized here
	User = &userRepository{session}
	User.CreateIndex()
}

func getSession() *mgo.Session {
	log.Println("Starting mongo session ", "Address: ", mongoPort)
	session, err := mgo.Dial(mongoPort)
	if err != nil {
		log.Fatalln("Couldn't connect to database,", "error:", err.Error(), "is database on?")
	}
	return session
}

func stringIDtoObjectID(id string) (bson.ObjectId, error) {
	log.Println("Converting string ID to mongoID")
	if !bson.IsObjectIdHex(id) {
		log.Println("ID is not mongo objectID, ", "ID: ", id)
		return "", errors.NewApiError(errors.InvalidID, fmt.Sprintf("Not a valid mongo ID: %s", id))
	}
	return bson.ObjectIdHex(id), nil
}

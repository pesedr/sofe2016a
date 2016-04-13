package repo

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const testDatabase = "test"

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("Couldn't connect to db", err)
	}

	return session
}

func stringIDtoObjectID(id string) (bson.ObjectId, error) {

	//TODO error handling
	if !bson.IsObjectIdHex(id) {
		fmt.Println("ID is not mongo objectID")
		return "", nil
	}
	return bson.ObjectIdHex(id), nil

}

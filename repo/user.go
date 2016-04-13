package repo

import (
	"fmt"

	"github.com/pesedr/sofe2016a/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	Get(userID string) (*models.User, error)
	Update(userID string, updatedUser *models.User) (*models.User, error)
	Delete(userID string) error
}

type userRepository struct {
	session *mgo.Session
}

var User UserRepository

const userCollection = "users"

func init() {
	User = &userRepository{getSession()}
}

func (u *userRepository) Create(user *models.User) (*models.User, error) {

	userCollection := u.collectionFromSession()
	err := userCollection.Insert(&user)
	if err != nil {
		fmt.Println("Couldn't insert object")
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Get(userID string) (*models.User, error) {
	user := &models.User{}

	if !bson.IsObjectIdHex(userID) {
		fmt.Println("ID is not mongo objectID")
		return nil, nil
	}
	oid := bson.ObjectIdHex(userID)

	userCollection := u.collectionFromSession()
	err := userCollection.FindId(oid).One(&user)
	if err != nil {
		fmt.Println("Couldn't find object", err)
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Update(userID string, updatedUser *models.User) (*models.User, error) {

	if !bson.IsObjectIdHex(userID) {
		fmt.Println("ID is not mongo objectID")
		return nil, nil
	}
	oid := bson.ObjectIdHex(userID)

	updatedUser.ID = oid

	userCollection := u.collectionFromSession()
	update := bson.M{"$set": updatedUser}
	err := userCollection.UpdateId(oid, update)
	if err != nil {
		fmt.Println("error updating user", err)
		return nil, err
	}
	return updatedUser, err
}

func (u *userRepository) Delete(userID string) error {
	return nil
}

// TODO figure out how to extract to a parent method for all db collection structs
func (u *userRepository) collectionFromSession() *mgo.Collection {
	return u.session.DB(testDatabase).C(userCollection)
}

package repo

import (
	"fmt"
	"log"

	"github.com/pesedr/sofe2016a/errors"
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

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	userCollection := u.collectionFromSession()

	log.Println("Inserting user into db")
	err := userCollection.Insert(&user)
	if err != nil {
		log.Println("Insert failed", "error:", err.Error())
		return nil, errors.NewApiError(errors.DatabaseError, fmt.Sprintf("error inserting user into DB, userID %s", user.ID))
	}

	return user, nil
}

func (u *userRepository) Get(userID string) (*models.User, error) {
	user := &models.User{}

	oid, err := stringIDtoObjectID(userID)
	if err != nil {
		return nil, err
	}

	userCollection := u.collectionFromSession()

	log.Println("Searching for user in DB")
	err = userCollection.FindId(oid).One(&user)
	if err != nil {
		log.Println("Could not find user", "error:", err.Error())
		return nil, errors.NewApiError(errors.UserNotFound, fmt.Sprintf("userID not found, id: %s", userID))
	}

	return user, nil
}

func (u *userRepository) Update(userID string, updatedUser *models.User) (*models.User, error) {
	oid, err := stringIDtoObjectID(userID)
	if err != nil {
		return nil, err
	}
	updatedUser.ID = oid

	userCollection := u.collectionFromSession()

	log.Println("Updating user in DB")
	update := bson.M{"$set": updatedUser}
	err = userCollection.UpdateId(oid, update)
	if err != nil {
		log.Println("Could not update user", "error:", err.Error())
		return nil, errors.NewApiError(errors.UserNotFound, fmt.Sprintf("user could not be updated, id: %s", userID))
	}
	return updatedUser, err
}

func (u *userRepository) Delete(userID string) error {
	return nil
}

func (u *userRepository) collectionFromSession() *mgo.Collection {
	return u.session.DB(testDatabase).C(userCollection)
}

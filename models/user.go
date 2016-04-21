package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

var UsersDB map[string]*User

func init() {
	UsersDB = make(map[string]*User)
}

type User struct {
	// ID        bson.ObjectId `json:"id" bson:"_id"`
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Account   *Accounts `json:"account"`
}

func (u *User) String() string {
	return fmt.Sprintf("userID: %s, name: %s, email: %s %s", u.ID, u.FirstName, u.LastName, u.Email)
}

type Account struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Balance float64       `json:"balance"`
}

type Accounts []Account

func (a *Account) Transfer() {}

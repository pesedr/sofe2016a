package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	// Account   *Accounts     `json:"account"`
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Password  string        `json:"password"`
	Username  string        `json:"username"`
}

func (u *User) String() string {
	return fmt.Sprintf("userID: %s, name: %s, email: %s %s", u.ID, u.FirstName, u.LastName, u.Email)
}

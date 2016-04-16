package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"-"`
	Account  *Accounts     `json:"account"`
}

type Account struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Balance float64       `json:"balance"`
}
type Accounts []Account

func (a *Account) Transfer() {}

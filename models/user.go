package models

var Users = map[int]*User{}
var Seq = 1

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

package models

import (
//"labix.org/v2/mgo"
//"labix.org/v2/mgo/bson"
)

type Account struct {
	Id       int
	Fname    string
	Lname    string
	Gradyear int
	Email    string
}

func FindAccount(id int) Account {
	return Account{1, "Bob", "Jones", 2018, "bjones@northwestern.edu"}
}

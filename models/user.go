package models

import (
	"fmt"
	"time"
)

const (
	USER_ID_LENGTH = 32
)

type UserID	[USER_ID_LENGTH]byte

type User interface {
	String() string // string method
	Name() sring // user name
	Identity()	UserID // return user identity
	Properties() []Property // list all user's properties
	GetProb(ItemID) Property // get a property by its identity
	AddProperty(Property) bool // add a property
}


type UserProfile struct {
	Name		string
	Location	string
	Birth		string	
}



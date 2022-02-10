package models

import (
	"time"
)

type DB interface {
	GetSession(SessionToken) (Session, error) //get a session by its id
	GetUser(UserID) (User, error) // get user from his id
	GetItem(ItemID) (Item, error)
	AddSession(Session) error
	AddUser(User) error
	AddItem(Item) error
	GetEvent(EventID) Event
	GetUserEvents(uid UserID, from time.Time, to time.Time, lim int) []EventID
	GetUserLatestEvents(uid UserID, lim int) []EventID
}

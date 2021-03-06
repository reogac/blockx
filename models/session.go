package models

import (
	"time"
)


type SessionToken string

type Session interface {
	GetToken() SessionToken
	GetUser() UserID
	LoginTime()	time.Time
	LastSeenTime() time.Time
}


type SessionManager	interface {
	Get(SessionToken) (Session, error) //get a session by its token
	Add(Session) error //add a new session
	Delete(SessionToken) error //delete a session
}

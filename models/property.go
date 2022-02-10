package models

import (
	"fmt"
	"time"
)

type Property interface {
	Item	
	GetEvents
	Owner() UserID
	Created() time.Time
}


type geneticProperty struct {
	owner			string //owner identity
	created			time.Time // time when it was created
	lastevent		time.Time	// time when last event updated
}


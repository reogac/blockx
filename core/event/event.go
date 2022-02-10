package event

import (
	"fmt"
	"time"
)
const (
	EVENT_ID_LENGTH	= 32
)

type EventWho 	string
type EventWhere	string
type EventWhat	string
type EventWhen	time.Time
type EventWhy	string
type EventID	[EVENT_ID_LENGTH]byte

type Event struct {
	who		EventWho
	where	EventWhere
	what	EventWhat
	when	EventWhen
	why		EventWhy
}


func (e *Event) Marshal() []byte {
	return []byte{}
}

func (* *Event) Unmarshal(wire []byte) error {
	return nil
}

// return id of the event (its hash value)
func (e *Event) GetID() (id EventID) {
	return
}

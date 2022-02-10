package models

import (
	"time"
)

const (
	EVENT_ID_LENGTH = 32
)

type EventID	[EVENT_ID_LENGTH]byte

type Event interface {
	Created() time.Time
	On() Property
	Owner()	User //who create the event
}


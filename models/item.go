package models

const (
	ITEM_ID_LENGTH = 32
)

type ItemID	[ITEM_ID_LENGTH]byte

type Item interface {
	GetID()
	String()
}


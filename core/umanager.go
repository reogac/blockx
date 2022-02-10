package core

import (
	"blockx/models"
)

type umImpl struct {
}


func NewUserManager() models.UserManager {
	return &umImpl{}
}

func (um *umImpl) Get(id models.UserID) (models.User, error) {
	return nil, nil
}


func (um *umImpl) Add(u models.User) error {
	return nil
}


func (um *umImpl) Delete(id models.UserID) error {
	return nil
}


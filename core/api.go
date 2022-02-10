package core

import (
	"blockx/models"
)

type UserApi interface {
	UserAddEvent()
}

type UserManagerApi interface {
	UserRegister()
	UserLogin()
	UserLogout()
}

type ViewerApi interface {
	GetOwner(models.ItemID) models.User //get owner of a product
}


type Api interface {
	UserApi
	UserManagerApi
}


type apiImpl struct {
}

// user login 
func (api apiImpl) UserLogin(c models.UserCredential) (models.User, error) {
	return nil, nil
}

// user log out, clean up resource
func (api apiImpl) UserLogout(u models.User) {
}

// user add an event
func (api apiImpl) UserAddEvent(u models.User, e models.Event) error {
	return nil
}

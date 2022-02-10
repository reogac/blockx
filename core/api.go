package core

type UserApi interface {
	UserAddEvent()
}

type UserManagerApi interface {
	UserRegister()
	UserLogin()
	UserLogout()
}

type ViewerApi interface {
	GetOwner(ItemID) User //get owner of a product
}


type Api interface {
	UserApi
	UserManagerApi
}


type apiImpl struct {
}

// user login 
func (api apiImpl) UserLogin(c UserCredential) (User, error) {
	return nil, nil
}

// user log out, clean up resource
func (api apiImpl) UserLogout(u User) {
}

// user add an event
func (api apiImpl) UserAddEvent(u User, e Event) error {
	return nil
}

package core


import (
	"blockx/models"
)

type App interface {
	Run() error
	Stop()
}


func NewApp() App {
	app := &appImpl{
		smanager:	NewSessionManager(),
		umanager:	NewUserManager(),
	}
	return app
}


type appImpl struct {
	smanager	models.SessionManager
	umanager	models.UserManager
}


func (a *appImpl) Run() error {
	return nil
}


func (a *appImpl) Stop() {
}

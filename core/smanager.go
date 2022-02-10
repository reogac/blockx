package core

import (
	"blockx/models"
)

type smImpl struct {
}

func NewSessionManager() models.SessionManager {
	return &smImpl{}
}

func (sm *smImpl) Get(token models.SessionToken) (models.Session, error) {
	return nil, nil
}


func (sm *smImpl) Add(s models.Session) error {
	return nil
}

func (sm * smImpl) Delete(token models.SessionToken) error {
	return nil
}

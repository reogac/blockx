package models


type Authenticator interface {
	Authenticate(UserCredential) error
}

package authentication

import (
	"errors"

	firebase "firebase.google.com/go"
)

type (
	// authPersistance is a persistance to preserve session management connection
	AuthPersistance struct {
		App *firebase.App
	}

	// AuthRepository defines session management methods available
	AuthRepository interface {
		CreateSessionCookie(*string)(error)
		DestroySessionCookie()(error)
		VerifySessionCookie()(error)
	}
)

// NewAuthPersistance creates a new auth persistance
func NewAuthPersistance() AuthRepository {
	return AuthPersistance{
		App: getFirebaseConnection(),
	}
}

// CreateSessionCookie creates and set session cookie from idToken
func (authPersistance AuthPersistance) CreateSessionCookie(idToken *string) error {
	return errors.New("failed to create session cookie")
}

// DestroySessionCookie destroys session cookie
func (authPersistance AuthPersistance) DestroySessionCookie() error {
	return errors.New("failed to create session cookie")
}

// VerifySessionCookie verifies session cookie attached to the request
func (authPersistance AuthPersistance) VerifySessionCookie() error {
	return errors.New("failed to create session cookie")
}
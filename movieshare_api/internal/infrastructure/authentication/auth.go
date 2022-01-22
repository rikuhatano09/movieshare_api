package authentication

import (
	"errors"
	"os"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
)

type (
	// NewAuthClient is a client for firebase sdk
	AuthClient struct {
		Client *auth.Client
	}

	// AuthRepository defines session management methods available
	AuthActions interface {
		CreateSessionCookie(*gin.Context)(error)
		DestroySessionCookie(*gin.Context)(error)
		VerifySessionCookie(*gin.Context)(error)
	}
)

// NewAuthClient creates a new auth client
func NewAuthClient(context *gin.Context) (AuthActions, error) {
	app := getFirebaseConnection()
	client, err := app.Auth(context)
	if err != nil {
		return AuthClient{}, err
	}
	return AuthClient{
		Client: client,
	}, nil
}

// CreateSessionCookie creates and set session cookie from idToken
func (authClient AuthClient) CreateSessionCookie(context *gin.Context) error {
	requestBody := contract.LoginPostRequestBody{}

	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		return err
	}
	
	// cookie expires in 30 minutes
	expiresIn := time.Minute * 30 

	// generate session cookie from IdToken 
	cookie, err := authClient.Client.SessionCookie(context, requestBody.IdToken, expiresIn)
	if err != nil {
		return err
	}

	// get the domain of the client
	stage := os.Getenv("GO_ENV")
	var domain string
	switch stage {
	case "local":
		domain = "localhost"
	case "prod":
		domain = ""
	}

	context.SetCookie(
		"session",
		cookie,
		int(expiresIn.Seconds()),
		"/",
		domain,
		true,
		true,
	)

	return nil
}

// DestroySessionCookie destroys session cookie
func (authClient AuthClient) DestroySessionCookie(context *gin.Context) error {
	return errors.New("failed to create session cookie")
}

// VerifySessionCookie verifies session cookie attached to the request
func (authClient AuthClient) VerifySessionCookie(context *gin.Context) error {
	return errors.New("failed to create session cookie")
}
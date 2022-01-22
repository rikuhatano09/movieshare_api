package usecase

import (
	"github.com/rikuhatano09/movieshare_api/internal/infrastructure/authentication"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
)

func CreateSessionCookie(requestBody contract.LoginPostRequestBody) error {
	authPersistance := authentication.NewAuthPersistance()

	err := authPersistance.CreateSessionCookie(&requestBody.IdToken)
	if err != nil {
		return err
	}
	return nil
}

func DestroySessionCookie() error {
	authPersistance := authentication.NewAuthPersistance()

	err := authPersistance.DestroySessionCookie()
	if err != nil {
		return err
	}
	return nil
}

func VerifySessionCookie() error {
	authPersistance := authentication.NewAuthPersistance()

	err := authPersistance.VerifySessionCookie()
	if err != nil {
		return err
	}
	return nil
}
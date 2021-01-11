package req

import (
	"errors"
	"mundipagg-sdk/internal/utils"
)

// Login A struct that holds the login information
type Login struct {
	BasicAuthUserName  string
	BasicAuthPassword  string
	BasicSecretAuthKey string
	Request            *Request
}

// NewMundipagg Create a Mundipagg object
func NewMundipagg(basicAuthUserName string, basicAuthPassword string, basicSecretAuthKey string) (*Login, error) {
	// Verify parameters
	if utils.IsStringEmpty(basicAuthUserName) {
		return nil, errors.New("Please provide basicAuthUserName")
	}

	if utils.IsStringEmpty(basicAuthPassword) {
		return nil, errors.New("Pleease provide basicAuthPassword")
	}

	if utils.IsStringEmpty(basicSecretAuthKey) {
		return nil, errors.New("Pleease provide basicSecretAuthKey")
	}

	// Create the mundipagg object
	login := &Login{
		BasicAuthUserName:  basicAuthUserName,
		BasicAuthPassword:  basicAuthPassword,
		BasicSecretAuthKey: basicSecretAuthKey,
	}

	return login, nil
}

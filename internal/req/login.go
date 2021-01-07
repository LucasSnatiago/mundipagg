package req

import (
	"errors"
	"mundipagg-sdk/internal/utils"
)

// Login A struct that holds the login information
type Login struct {
	BasicAuthUserName  string `json:"basicAuthUserName"`
	BasicAuthPassword  string `json:"basicAuthPassword"`
	BasicSecretAuthKey string
	// Define if the API is ready for production or for testing
	IsProduction bool
}

// NewMundipagg Create a Mundipagg object
func NewMundipagg(basicAuthUserName string, basicAuthPassword string, basicSecretAuthKey string, isProduction bool) (*Login, error) {
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
		IsProduction:       isProduction,
	}

	return login, nil
}

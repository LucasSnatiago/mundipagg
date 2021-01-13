package req

import (
	"errors"

	"github.com/lusantisuper/mundipagg/internal/utils"
)

// Login A struct that holds the login information
type Login struct {
	BasicSecretAuthKey string
	Request            *Request
}

// NewMundipagg Create a Mundipagg object
func NewMundipagg(basicSecretAuthKey string) (*Login, error) {
	// Verify parameters
	if utils.IsStringEmpty(basicSecretAuthKey) {
		return nil, errors.New("Pleease provide basicSecretAuthKey")
	}

	// Create the mundipagg object
	login := &Login{
		BasicSecretAuthKey: basicSecretAuthKey,
	}

	return login, nil
}

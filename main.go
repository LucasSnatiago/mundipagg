package main

import (
	"mundipagg-sdk/internal/req"
)

// NewMundipagg Create a Mundipagg object
func NewMundipagg(basicAuthUserName string, basicAuthPassword string, basicSecretAuthKey string, isProduction bool) (*req.Login, error) {
	return req.NewMundipagg(basicAuthUserName, basicAuthPassword, basicSecretAuthKey, isProduction)
}

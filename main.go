package main

import (
	"fmt"
	"mundipagg-sdk/internal/req"
	"mundipagg-sdk/mundipagg"
)

// NewMundipagg Create a Mundipagg object
func NewMundipagg(basicAuthUserName string, basicAuthPassword string, basicSecretAuthKey string) (*req.Login, error) {
	return req.NewMundipagg(basicAuthUserName, basicAuthPassword, basicSecretAuthKey)
}

func main() {
	mund, err := NewMundipagg("acc_eOlrOxofM2TAjxnB", "pk_test_yWBnrNqc4YFjanQO", SECRETKEY)
	if err != nil {
		panic(err)
	}

	subs := mundipagg.NewSubscription()
	subs.CustomerID = "cus_pL3yKGEsviMQo9Nz"
	subs.MinimumPrice = 1000
	subs.Description = "TESTE AVUL"
	subs.BillingTypes(1)
	subs.Currencys(1)
	subs.Intervals(3)
	subs.PaymentMethods(3)

	setup := mundipagg.Setup{
		Amount:      10000,
		Description: "Preco instalar",
		Payment:     &mundipagg.Payment{},
	}
	setup.Payment.PaymentMethods(2)
	subs.Setup = &setup

	mund.Request = &req.Request{
		Data: subs,
	}

	resp := mund.MakePostRequest()
	fmt.Println(resp)
}

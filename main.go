package main

import (
	"fmt"

	"github.com/lusantisuper/Mundipagg-SDK-GO/internal/req"
	"github.com/lusantisuper/Mundipagg-SDK-GO/mundipagg"
)

// NewMundipagg Create a Mundipagg object
func NewMundipagg(basicAuthUserName string, basicAuthPassword string, basicSecretAuthKey string) (*req.Login, error) {
	return req.NewMundipagg(basicAuthUserName, basicAuthPassword, basicSecretAuthKey)
}

func main() {
	mund, err := NewMundipagg(ACCOUNTSECRET, PASSWORDSECRET, SECRETKEY)
	if err != nil {
		panic(err)
	}

	subs := mundipagg.NewSubscription()
	subs.CustomerID = "cus_pL3yKGEsviMQo9Nz"
	subs.MinimumPrice = 1000
	subs.Description = "TESTE AVULSO"
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

	subs.Setup.Payment.PaymentMethods(2)

	boleto := &mundipagg.Boleto{
		Instructions: "Pague o boleto!",
		NossoNumero:  "123",
	}
	boleto.BankTypes(1)
	boleto.BoletoTypes(1)

	subs.Setup.Payment.Boleto = boleto

	schema := &mundipagg.PriceSchema{
		Price: 1000,
	}
	schema.SchemaTypes(1)

	item := mundipagg.Item{
		Name:          "Plano 1",
		PricingSchema: schema,
		Description:   "Este é o plano 1",
		Quantity:      1,
	}

	items := make([]mundipagg.Item, 1)
	items[0] = item
	subs.Items = &items

	mund.Request = &req.Request{
		Data: subs,
	}

	resp := mund.MakePostRequest()
	fmt.Println(resp)
}

package mundipagg

// Subscription A structs to create a json to subscription
type Subscription struct {
	Code          string `json:"code,omitempty"`
	PaymentMethod string
}

/*
paymentMethods := map[int]string {
	1: "credit_card",
	2: "boleto",
	3: "debit_card",
}
*/

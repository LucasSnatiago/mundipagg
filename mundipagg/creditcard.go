package mundipagg

// CreditCard Struct that holds a Credit Card object
type CreditCard struct {
	Installments int32 `json:"installments,omitempty"`
	// StatementDescriptor description that appears on the credit card bill -> Max 22 characters
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// OperationType if the transactions needs to be captuured
	OperationType string `json:"Operation_type,omitempty"`

	// Credit cards choose one
	CardCredit *CreditCard `json:"credit_card,omitempty"`
	CardID     string      `json:"card_id,omitempty"`
	CardToken  interface{} /// TODO ------------------------------------------

	// Recurrence define if this payment repeats or not (default not)
	Recurrence bool `json:"recurrence,omitempty"`

	// Metadata Extra information about the payment
	Metadata *interface{} // TODO ------------------------------------------

	// Only for private label cards
	ExtendedLimitEnabled bool `json:"extended_limit_enabled,omitempty"`
	// Code for the super limit
	ExtendedLimitCode string `json:"extended_limit_code,omitempty"`
	// Code for the type of store
	MerchantCategoryCode int32 `json:"merchant_id,omitempty"`

	// Authentication
	Authentication *interface{} // TODO ----------------------
	// AutoRecovery from a offline fail
	AutoRecovery bool `json:"auto_recovery,omitempty"`

	// Payload for google pay, apple pay and samsung pay
	Payload *interface{} `json:"payload,omitempty"` // TODO ---------------
}

/* OperationTypes
1 - auth_and_capture (default)
2 - auth_only
3 - pre_auth
*/
func (c CreditCard) OperationTypes(operationType int) string {
	operation := map[int]string{
		1: "auth_and_capture",
		2: "auth_only",
		3: "pre_auth",
	}

	return operation[operationType]
}

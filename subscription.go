package mundipagg

import "time"

// NewSubscription Return a New Subscription
func (m mundipagg) NewSubscription(s *Subscription, indepotencyKey string) (*Response, error) {

	resp, err := MakePostRequest(s, m.BasicSecretAuthKey, indepotencyKey, SUBSCRIPTIONURL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Subscription A structs to create a json to subscription
type Subscription struct {
	// Stores code
	Code string `json:"code,omitempty"`

	// Setting the payment
	PaymentMethod string `json:"payment_method,omitempty"`
	Currency      string `json:"currency,omitempty"`

	// Leave it empty if you want to start the subscription now
	StartAt *time.Time `json:"startAt,omitempty"`

	// Minimum value in cents
	MinimumPrice int32 `json:"minimum_price,omitempty"`

	// Interval between the charges
	Interval      string `json:"interval,omitempty"`
	IntervalCount int32  `json:"interval_count,omitempty"`
	BillingType   string `json:"billing_type,omitempty"`
	// If BillingType is a exact day set this variable
	BillingDay int32 `json:"billing_day,omitempty"`

	// Description for the payment
	Description string `json:"description,omitempty"`

	// CREDIT CARD
	// If the payments is credit card set the number of installments
	Installments int32 `json:"installments,omitempty"`
	// Text to appear on the credit card bill
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// CUSTOMER
	// The ID of the customer, needed if the customer variable is empty
	CustomerID string `json:"customer_id,omitempty"`
	// Customer, needed if the CustomerID is empty
	Customer *Customer `json:"customer,omitempty"`

	// Bill extras
	// Discounts
	Discounts *[]BillExtras `json:"discount,omitempty"`
	// Increments
	Increments *[]BillExtras `json:"increments,omitempty"`

	// Subscription items
	Items *[]Item `json:"items,omitempty"`

	// Price for the first payment (the setup)
	Setup *Setup `json:"setup,omitempty"`

	// GatewayAffiliationID affiliation id gived by the store
	GatewayAffiliationID string `json:"gateway_affiliation_id,omitempty"`
	// BoletoDueDays Days to expire the boleto, not passing
	BoletoDueDays int32 `json:"boleto_due_days,omitempty"`

	// Card to make the credit or debit payment
	Card *SubscriptionCards `json:"card,omitempty"`
	// Boleto to make a boleto payment
	Boleto *Boleto `json:"boleto,omitempty"`

	// Metadata extra information to the subscription
	Metadata interface{} `json:"metadata,omitempty"`
}

// SubscriptionCards necessary card information to make a subscription payment
type SubscriptionCards struct {
	// CardID pass the mundipagg ID for the card
	CardID string `json:"card_id,omitempty"`
	// Token pass the mundipagg token for the card
	Token string `json:"token,omitempty"`
}

/* PaymentMethods
1 - For credit payment
2 - For Debit Payment
3 - For Boleto
*/
func (s *Subscription) PaymentMethods(paymentType int) {
	paymentMethods := map[int]string{
		1: "credit_card",
		2: "debit_card",
		3: "boleto",
	}

	s.PaymentMethod = paymentMethods[paymentType]
}

/* Currencys
1 - BRL
2 - ARS
3 - BOB
4 - CLP
5 - COP
6 - MXN
7 - PYG
8 - USD
9 - UYU
10 - EUR
*/
func (s *Subscription) Currencys(currencyType int) {
	currencys := map[int]string{
		1:  "BRL",
		2:  "ARS",
		3:  "BOB",
		4:  "CLP",
		5:  "COP",
		6:  "MXN",
		7:  "PYG",
		8:  "USD",
		9:  "UYU",
		10: "EUR",
	}

	s.Currency = currencys[currencyType]
}

/* Intervals
1 - day
2 - week
3 - month
4 - year
*/
func (s *Subscription) Intervals(intervalType int) {
	interval := map[int]string{
		1: "day",
		2: "week",
		3: "month",
		4: "year",
	}

	s.Interval = interval[intervalType]
}

/* BillingTypes
1 - prepaid
2 - postpaid
3 - exact_day
*/
func (s *Subscription) BillingTypes(billingType int) {
	billing := map[int]string{
		1: "prepaid",
		2: "postpaid",
		3: "exact_day",
	}

	s.BillingType = billing[billingType]
}

// BillExtras extra information about discount and increments in the payments
type BillExtras struct {
	ID           string `json:"id,omitempty"`
	Cycles       int32  `json:"cycles"`
	Value        int32  `json:"value"`
	DiscountType string `json:"discount"`
	// Write the id if you want to inform the product to recive the discount or increment
	// Leave empty if you want to apply to the entire bill
	ItemID    string     `json:"item_id,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

/* DiscountTypes
1 - flat
2 - percentage
*/
func (bs *BillExtras) DiscountTypes(discountType int) {
	discount := map[int]string{
		1: "flat",
		2: "percentage",
	}

	bs.DiscountType = discount[discountType]
}

/* StatusTypes
1 - active
2 - deleted
*/
func (bs *BillExtras) StatusTypes(statusType int) {
	status := map[int]string{
		1: "active",
		2: "deleted",
	}

	bs.Status = status[statusType]
}

// Item to the subscription
type Item struct {
	// ID starts with "si_"
	ID string `json:"id,omitempty"`
	// Description for the item
	Description string `json:"description,omitempty"`
	// Number of times to be charged
	Cycles int32 `json:"cycles,omitempty"`
	// Quantity of itens
	Quantity int32 `json:"quantity,omitempty"`

	// Item status
	Status string `json:"status,omitempty"`

	// PricingSchema schema of the pricing
	PricingSchema *PriceSchema `json:"pricing_scheme,omitempty"`

	// Created / Updated / Deleted
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// Discounts and increments for this item
	Discounts  *[]BillExtras `json:"discounts,omitempty"`
	Increments *[]BillExtras `json:"increments,omitempty"`

	// Item name
	Name string `json:"name,omitempty"`
}

/* StatusTypes
1 - active
2 - inactive
3 - deleted
*/
func (i *Item) StatusTypes(statusType int) {
	status := map[int]string{
		1: "active",
		2: "inactive",
		3: "deleted",
	}

	i.Status = status[statusType]
}

// Setup for the payments
type Setup struct {
	Amount      int32    `json:"amount"`
	Description string   `json:"description"`
	Payment     *Payment `json:"payment"`
}

// Payment struct for a payment
type Payment struct {
	// PaymentMethod Choose the payment method
	PaymentMethod string `json:"payment_method,omitempty"`
	// CreditCard if want credit card, add it here
	CreditCard *CreditCard `json:"credit_card,omitempty"`
	// Voucher
	Voucher interface{} `json:"voucher,omitempty"`
	// Boleto if wants boleto, add it here
	Boleto       *Boleto     `json:"boleto,omitempty"`
	BankTransfer interface{} `json:"bank_transfer,omitempty"`
	Checkout     interface{} `json:"checkout,omitempty"`
	Cash         interface{} `json:"cash,omitempty"`
	// Amount the price in cents
	Amount int64 `json:"amount,omitempty"`
	// Metadata adds the extra information here
	Metadata interface{} `json:"metadata,omitempty"`
	// GatewayAffiliationID your gateway affiliation ID
	GatewayAffiliationID string `json:"gateway_affiliation_id,omitempty"`
}

/* PaymentMethods definitions
1 - credit_card
2 - boleto
3 - voucher TODO
4 - bank_transfer TODO
5 - safety_pay TODO
6 - checkout TODO
7 - cash TODO
*/
func (p *Payment) PaymentMethods(paymentType int) {
	payment := map[int]string{
		1: "credit_card",
		2: "boleto",
	}

	p.PaymentMethod = payment[paymentType]
}

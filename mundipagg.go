package mundipagg

// Mundipagg struct that holds all functions working in the lib
type Mundipagg interface {
	NewSubscription(s *Subscription, indepotencyKey string) (*Response, error)
	NewCustomer(c *Customer, indepotencyKey string) (*Response, error)
	NewCardByToken(customerID string, cardToken string, indepotencyKey string) (*Response, error)
	UpdateStartAt(u *UpdateStartAtSubscription, mundipaggCustomerID string, indepotencyKey string) (*Response, error)
	UpdateNextBillingDay(u *UpdateNextBillingDaySubscription, mundipaggCustomerID string, indepotencyKey string) (*Response, error)
	AddDiscount(b *BillExtras, mundipaggSubscriptionID string, indepotencyKey string) (*Response, error)
}

// Login A struct that holds the login information
type mundipagg struct {
	BasicSecretAuthKey string
}

// New Create a Mundipagg object
func New(basicSecretAuthKey string, testBasicSecretKey string, isProduction bool) (Mundipagg, error) {
	secret := basicSecretAuthKey
	if !isProduction {
		secret = testBasicSecretKey
	}

	// Create the mundipagg object
	mund := mundipagg{
		BasicSecretAuthKey: secret,
	}

	return mund, nil
}

func main() {}

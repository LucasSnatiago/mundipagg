package mundipagg

import "time"

// UpdateStartAtSubscription update the start at subscription
type UpdateStartAtSubscription struct {
	StartAt *string `json:"start_at,omitempty"`
}

// UpdateStartAt make the request to mundipagg updating the start at return the ID of the object
func (m mundipagg) UpdateStartAt(u *UpdateStartAtSubscription, mundipaggCustomerID string, indepotencyKey string) (*Response, error) {
	completeURL := SUBSCRIPTIONURL + "/" + mundipaggCustomerID + SUBSCRIPTIONUPDATESTARTATURL

	resp, err := MakePostRequest(u, m.BasicSecretAuthKey, indepotencyKey, completeURL, "PATCH")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateNextBillingDaySubscription update the next billing day in the subscription
type UpdateNextBillingDaySubscription struct {
	NextBillingDay *time.Time `json:"next_billing_at,omitempty"`
}

// UpdateNextBillingDay make the request to mundipagg updating the next billing day return the ID of the object
func (m mundipagg) UpdateNextBillingDay(u *UpdateNextBillingDaySubscription, mundipaggCustomerID string, indepotencyKey string) (*Response, error) {
	completeURL := SUBSCRIPTIONURL + "/" + mundipaggCustomerID + SUBSCRIPTIONUPDATENEXTBILLINGDAYURL

	resp, err := MakePostRequest(u, m.BasicSecretAuthKey, indepotencyKey, completeURL, "PATCH")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

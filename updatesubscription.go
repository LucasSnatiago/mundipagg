package mundipagg

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

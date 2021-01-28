package mundipagg

func (m mundipagg) AddDiscount(b *BillExtras, mundipaggSubscriptionID string, indepotencyKey string) (*Response, error) {
	completeURL := SUBSCRIPTIONURL + "/" + mundipaggSubscriptionID + DISCOUNTENDPOINT

	resp, err := MakePostRequest(b, m.BasicSecretAuthKey, indepotencyKey, completeURL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

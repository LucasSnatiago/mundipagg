package mundipagg

// BASEURL Url for to send information to the API
const BASEURL string = "https://api.mundipagg.com/core/v1/"

// CUSTOMERURL Link to get from api a customer
const CUSTOMERURL string = BASEURL + "customers"

// SUBSCRIPTIONURL Link to make a subscription
const SUBSCRIPTIONURL string = BASEURL + "subscriptions"

// SUBSCRIPTIONUPDATENEXTBILLINGDAYURL end of the url for the billing date
const SUBSCRIPTIONUPDATENEXTBILLINGDAYURL string = "/billing-date"

// SUBSCRIPTIONUPDATESTARTATURL end of the subscription start at url
const SUBSCRIPTIONUPDATESTARTATURL string = "/start-at"

// CARDENDPOINT Links the customer to the credit card (ex: CUSTOMERURL + customer_id + "/cards")
const CARDENDPOINT string = "cards"

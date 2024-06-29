package models

type Billing struct {
	ID             int     `json:"id"`
	SubscriptionID int     `json:"subscription_id"`
	Amount         float64 `json:"amount"`
	BillingDate    string  `json:"billing_date"`
	Status         string  `json:"status"`
}

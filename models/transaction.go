package models

type Transaction struct {
	ID              int     `json:"id"`
	UserID          int     `json:"user_id"`
	SubscriptionID  int     `json:"subscription_id"`
	Amount          float64 `json:"amount"`
	TransactionDate string  `json:"transaction_date"`
	Description     string  `json:"description"`
}

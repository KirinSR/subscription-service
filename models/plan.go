package models

type Plan struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Price        float64    `json:"price"`
	BillingCycle string     `json:"billing_cycle"`
	MemberLimit  int        `json:"member_limit"`
	CreatedAt    string     `json:"created_at"`
	Features     []*Feature `json:"features,omitempty"`
}

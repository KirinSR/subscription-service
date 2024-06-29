package models

type Subscription struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	PlanID       int    `json:"plan_id"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	TrialEndDate string `json:"trial_end_date"`
	User         *User  `json:"user,omitempty"`
	Plan         *Plan  `json:"plan,omitempty"`
}

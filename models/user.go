package models

type User struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"-"`
	CompanyID int      `json:"company_id"`
	CreatedAt string   `json:"created_at"`
	Company   *Company `json:"company,omitempty"`
}

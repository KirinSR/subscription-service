package models

type Company struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	Users     []*User `json:"users,omitempty"`
}

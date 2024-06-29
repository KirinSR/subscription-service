package models

import (
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Price        float64    `json:"price"`
	BillingCycle string     `json:"billing_cycle"`
	MemberLimit  int        `json:"member_limit"`
	CreatedAt    time.Time  `json:"created_at"`
	CreateBy     uuid.UUID  `json:"create_by" gorm:"type:uuid;default:uuid_generate_v4()"`
	UpdatedAt    time.Time  `json:"updated_at"`
	UpdatedBy    uuid.UUID  `json:"updated_by" gorm:"type:uuid;default:uuid_generate_v4()"`
	Features     []*Feature `json:"features,omitempty"`
}

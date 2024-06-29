package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	PlanID       uuid.UUID `json:"plan_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	StartDate    string    `json:"start_date"`
	EndDate      string    `json:"end_date"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	CreateBy     uuid.UUID `json:"create_by" gorm:"type:uuid;default:uuid_generate_v4()"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    uuid.UUID `json:"updated_by" gorm:"type:uuid;default:uuid_generate_v4()"`
	TrialEndDate string    `json:"trial_end_date"`
	User         *User     `json:"user,omitempty"`
	Plan         *Plan     `json:"plan,omitempty"`
}

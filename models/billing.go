package models

import (
	"time"

	"github.com/google/uuid"
)

type Billing struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	SubscriptionID uuid.UUID `json:"subscription_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Amount         float64   `json:"amount"`
	BillingDate    string    `json:"billing_date"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	CreateBy       uuid.UUID `json:"create_by" gorm:"type:uuid;default:uuid_generate_v4()"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      uuid.UUID `json:"updated_by" gorm:"type:uuid;default:uuid_generate_v4()"`
}

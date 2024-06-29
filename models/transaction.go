package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID          uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	SubscriptionID  int       `json:"subscription_id"`
	Amount          float64   `json:"amount"`
	TransactionDate string    `json:"transaction_date"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	CreateBy        uuid.UUID `json:"create_by" gorm:"type:uuid;default:uuid_generate_v4()"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedBy       uuid.UUID `json:"updated_by" gorm:"type:uuid;default:uuid_generate_v4()"`
}

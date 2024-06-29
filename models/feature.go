package models

import (
	"time"

	"github.com/google/uuid"
)

type Feature struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreateBy    uuid.UUID `json:"create_by" gorm:"type:uuid;default:uuid_generate_v4()"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   uuid.UUID `json:"updated_by" gorm:"type:uuid;default:uuid_generate_v4()"`
}

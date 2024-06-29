package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CompanyID int       `json:"company_id"`
	Company   *Company  `json:"company,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	CreateBy  uuid.UUID `json:"create_by" gorm:"type:uuid;default:uuid_generate_v4()"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by" gorm:"type:uuid;default:uuid_generate_v4()"`
}

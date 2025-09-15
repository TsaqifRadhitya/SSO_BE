package Entity

import (
	"gorm.io/gorm"
	"time"
)

type VerifyToken struct {
	gorm.Model
	Token          string    `gorm:"type:text ;unique; not null;"`
	UserId         int       `gorm:"not null;"`
	IsUsed         bool      `gorm:"not null;default false;"`
	ApplicationKey string    `gorm:"type:text ; not null;"`
	ExpiresAt      time.Time `gorm:"not null;"`
	User           User      `gorm:"foreignkey:UserId"`
}

func (s *VerifyToken) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ExpiresAt.IsZero() {
		s.ExpiresAt = time.Now().Add(5 * time.Minute)
	}
	return
}

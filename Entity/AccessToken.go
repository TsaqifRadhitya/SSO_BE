package Entity

import (
	"gorm.io/gorm"
	"time"
)

type AccessToken struct {
	gorm.Model
	Token     string    `gorm:"not null;text"`
	ClientId  *int      `gorm:"not null"`
	Client    *Client   `gorm:"foreignkey:ClientId" json:"-"`
	ExpiresAt time.Time `gorm:"not null"`
}

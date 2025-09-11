package Entity

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	JwtToken      string    `gorm:"not null; text" json:"jwt_token"`
	JwtExpiry     time.Time `gorm:"not null;" json:"-"`
	RefreshToken  string    `gorm:"not null; text" json:"refresh_token"`
	RefreshExpiry time.Time `gorm:"not null; text" json:"-"`
	UserId        int       `gorm:"not null; text" json:"-"`
	User          User      `gorm:"foreignkey:UserId" json:"-"`
}

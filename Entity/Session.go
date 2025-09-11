package Entity

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	JwtToken      string    `gorm:" text; not null;" json:"jwt_token"`
	JwtExpiry     time.Time `gorm:"; not null;" json:"-"`
	RefreshToken  string    `gorm:"type:text; not null;" json:"refresh_token"`
	RefreshExpiry time.Time `gorm:"type:text; not null;" json:"-"`
	UserId        int       `gorm:"type:text; not null;" json:"-"`
	User          User      `gorm:"foreignkey:UserId" json:"-"`
}

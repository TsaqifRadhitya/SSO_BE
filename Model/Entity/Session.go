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

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	if s.JwtExpiry.IsZero() {
		s.JwtExpiry = time.Now().Add(15 * time.Minute)
	}

	if s.RefreshExpiry.IsZero() {
		s.RefreshExpiry = time.Now().AddDate(0, 1, 0)
	}
	return
}

func (s *Session) BeforeUpdate(tx *gorm.DB) (err error) {
	if s.JwtExpiry.IsZero() {
		s.JwtExpiry = time.Now().Add(15 * time.Minute)
	}

	if s.RefreshExpiry.IsZero() {
		s.RefreshExpiry = time.Now().AddDate(0, 1, 0)
	}
	return
}

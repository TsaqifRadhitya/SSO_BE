package Entity

import "gorm.io/gorm"

type VerifyToken struct {
	gorm.Model
	Token  string `gorm:"type:text ;unique; not null;"`
	UserId int    `gorm:"not null;"`
	User   User   `gorm:"foreignkey:UserId"`
}

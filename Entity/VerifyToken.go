package Entity

import "gorm.io/gorm"

type VerifyToken struct {
	gorm.Model
	Token  string `gorm:"unique; not null; text"`
	UserId int    `gorm:"not null;"`
	User   User   `gorm:"foreignkey:UserId"`
}

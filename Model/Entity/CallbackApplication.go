package Entity

import "gorm.io/gorm"

type CallbackApplication struct {
	gorm.Model
	Callback      string      `gorm:"type:varchar(255);not null"`
	ApplicationId int         `gorm:"not null"`
	Application   Application `gorm:"foreignkey:ApplicationId" json:"-"`
}

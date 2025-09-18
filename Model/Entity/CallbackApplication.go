package Entity

import "gorm.io/gorm"

type CallbackApplication struct {
	gorm.Model
	Callback      string      `gorm:"type:varchar(255);not null" json:"callback"`
	ApplicationId int         `gorm:"not null" json:"application_id"`
	Application   Application `gorm:"foreignkey:ApplicationId" json:"-"`
}

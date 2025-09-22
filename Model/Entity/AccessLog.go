package Entity

import "gorm.io/gorm"

type AccessLog struct {
	gorm.Model
	UserId        int         `gorm:"not null" json:"-"`
	User          User        `gorm:"foreignkey:UserId"`
	ApplicationId int         `gorm:"not null" json:"-"`
	Application   Application `gorm:"foreignkey:ApplicationId"`
}

package Entity

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ClientName          string                `gorm:"type:varchar(255); not null;"`
	Key                 string                `gorm:"type:text; not null;"`
	Callback            string                `gorm:"type:varchar(255); not null"`
	CallbackApplication []CallbackApplication `gorm:"foreignkey:ApplicationId;"`
}

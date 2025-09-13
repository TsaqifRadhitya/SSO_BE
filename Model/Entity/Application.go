package Entity

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	ClientName          string                `gorm:"type:varchar(255); not null;"`
	Key                 string                `gorm:"type:text; not null;"`
	ClientKey           string                `gorm:"type:text; not null;"`
	Callback            string                `gorm:"type:varchar(255); not null"`
	CallbackApplication []CallbackApplication `gorm:"foreignkey:ApplicationId;"`
	OwnerId             int                   `gorm:"not null;"`
	Owner               User                  `gorm:"foreignkey:OwnerId;" json:"-"`
}

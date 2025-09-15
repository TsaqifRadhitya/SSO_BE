package Entity

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	ApplicationName     string                `gorm:"type:varchar(255); not null;" json:"application_name"`
	ApplicationKey      string                `gorm:"type:text; not null;" json:"application_key"`
	CallbackApplication []CallbackApplication `gorm:"foreignkey:ApplicationId;" json:"callback_application,omitempty"`
	OwnerId             int                   `gorm:"not null;" json:"-"`
	Owner               User                  `gorm:"foreignkey:OwnerId;" json:"-"`
}

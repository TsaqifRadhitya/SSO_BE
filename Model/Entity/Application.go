package Entity

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	ApplicationName     string                `gorm:"type:varchar(255); not null;"`
	ApplicationKey      string                `gorm:"type:text; not null;"`
	CallbackApplication []CallbackApplication `gorm:"foreignkey:ApplicationId;"`
	OwnerId             int                   `gorm:"not null;"`
	Owner               User                  `gorm:"foreignkey:OwnerId;" json:"-"`
}

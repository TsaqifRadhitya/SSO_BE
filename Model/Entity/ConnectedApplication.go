package Entity

import "gorm.io/gorm"

type ConnectedApplication struct {
	gorm.Model
	ApplicationId int
	Application   Application `gorm:"foreignKey:ApplicationId;references:ID"`
	UserId        int
	User          User `gorm:"foreignKey:UserId;references:ID"`
}

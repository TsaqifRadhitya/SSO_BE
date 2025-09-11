package Entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Name     string `gorm:"type:varchar(255);not null"`
	Phone    string `gorm:"type:varchar(255);unique;not null"`
}

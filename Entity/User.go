package Entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"varchar(255); unique; not null"`
	Password string `gorm:"varchar(255); not null" json:"-"`
	Name     string `gorm:"varchar(255); not null"`
	Phone    string `gorm:"varchar(255); unique; not null"`
}

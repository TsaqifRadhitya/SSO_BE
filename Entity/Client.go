package Entity

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ClientName string `gorm:"not null;varchar(255)"`
	Key        string `gorm:"not null;text"`
	Callback   string `gorm:"varchar(255); not null"`
}

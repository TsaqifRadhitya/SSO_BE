package Config

import (
	Entity2 "SSO_BE_API/Model/Entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT,
	)
	if DB_PASS == "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=disable",
			DB_HOST, DB_USER, DB_NAME, DB_PORT,
		)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	if err = db.AutoMigrate(
		&Entity2.User{},
		&Entity2.Application{},
		&Entity2.CallbackApplication{},
		&Entity2.VerifyToken{},
		&Entity2.Session{}); err != nil {
		return err
	}
	return nil
}

func DbClose() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

package Config

import (
	Entity2 "SSO_BE_API/Model/Entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

var DB *gorm.DB

func createDSN() string {
	parts := []string{
		fmt.Sprintf("host=%s", DB_HOST),
		fmt.Sprintf("user=%s", DB_USER),
		fmt.Sprintf("dbname=%s", DB_NAME),
		fmt.Sprintf("port=%s", DB_PORT),
		fmt.Sprintf("sslmode=%s", SSL_MODE),
	}

	if DB_PASS != "" {
		parts = append(parts, fmt.Sprintf("password=%s", DB_PASS))
	}

	return strings.Join(parts, " ")
}

func DbConnect() error {
	dsn := createDSN()
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

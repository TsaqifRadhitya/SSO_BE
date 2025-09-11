package Config

import (
	"SSO_BE_API/Entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() error {
	pass := DB_PASS
	if pass == "" {
		pass = ""
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT,
	)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	DB.AutoMigrate(&Entity.User{})
	DB.AutoMigrate(&Entity.Client{})
	DB.AutoMigrate(&Entity.AccessToken{})
	DB.AutoMigrate(&Entity.VerifyToken{})
	DB.AutoMigrate(&Entity.Session{})
	return nil
}

func DbClose() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

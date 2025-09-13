package Config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	PORT                  string
	DB_USER               string
	DB_PASS               string
	DB_NAME               string
	DB_PORT               string
	DB_HOST               string
	JWT_KEY               string
	JWT_REFRESH_TOKEN_KEY string
	VERIFY_TOKEN_KEY      string
	ENV                   string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	PORT = os.Getenv("PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	JWT_KEY = os.Getenv("JWT_KEY")
	JWT_REFRESH_TOKEN_KEY = os.Getenv("JWT_REFRESH_TOKEN_KEY")
	VERIFY_TOKEN_KEY = os.Getenv("VERIFY_TOKEN_KEY")
	ENV = os.Getenv("ENV")
}

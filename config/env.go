package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                  string
	BaseURL               string
	Mode                  string
	JwtSecret             string
	JwtRefreshTokenSecret string

	MailerPort     string
	MailerHost     string
	MailerEmail    string
	MailerPassword string

	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string

	Debug string
}

var (
	envConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}

		envConfig = &Config{
			Port:                  os.Getenv("APP_PORT"),
			BaseURL:               os.Getenv("BASE_URL"),
			JwtSecret:             os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
			JwtRefreshTokenSecret: os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
			Mode:                  os.Getenv("MODE"),

			MailerPort:     os.Getenv("MAILER_PORT"),
			MailerHost:     os.Getenv("MAILER_HOST"),
			MailerEmail:    os.Getenv("MAILER_EMAIL"),
			MailerPassword: os.Getenv("MAILER_PASSWORD"),

			DBPort:     os.Getenv("DB_PORT"),
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBHost:     os.Getenv("DB_HOST"),

			Debug: os.Getenv("DEBUG"),
		}
	})

	return envConfig
}

package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	BaseURL string
	Mode    string
	Debug   string
	// Origins        string
	AllowedOrigins []string

	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string

	MailerPort     string
	MailerHost     string
	MailerEmail    string
	MailerPassword string

	JwtSecret             string
	JwtRefreshTokenSecret string
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

		// origins := os.Getenv("ALLOWED_ORIGINS")
		// var allowedOrigins []string

		// if origins != "" {
		// 	for _, origin := range strings.Split(origins, ",") {
		// 		allowedOrigins = append(allowedOrigins, strings.TrimSpace(origin))
		// 	}
		// }

		envConfig = &Config{
			Port:    os.Getenv("APP_PORT"),
			BaseURL: os.Getenv("BASE_URL"),
			Mode:    os.Getenv("MODE"),
			Debug:   os.Getenv("DEBUG"),
			// Origins:        origins,
			AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),

			DBPort:     os.Getenv("DB_PORT"),
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBHost:     os.Getenv("DB_HOST"),

			MailerPort:     os.Getenv("MAILER_PORT"),
			MailerHost:     os.Getenv("MAILER_HOST"),
			MailerEmail:    os.Getenv("MAILER_EMAIL"),
			MailerPassword: os.Getenv("MAILER_PASSWORD"),

			JwtSecret:             os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
			JwtRefreshTokenSecret: os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
		}
	})

	return envConfig
}

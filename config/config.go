package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User        string
	Pass        string
	Schema      string
	Host        string
	Port        string
	Database    string
	NameService string
	PortService string
	KeyOpenIa   string
}

func NewInitConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando archivo .env: %v", err)
	}
	return Config{
		User:        os.Getenv("DB_USER"),
		Pass:        os.Getenv("DB_PASS"),
		Schema:      os.Getenv("DB_SCHEMA"),
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		Database:    os.Getenv("DB_DATABASE"),
		NameService: os.Getenv("NAME_SERVICE"),
		PortService: os.Getenv("PORT_SERVICE"),
		KeyOpenIa:   os.Getenv("KEY_OPEN_IA"),
	}
}

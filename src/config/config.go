package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUri string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".env")
	return &Config{
		PostgresUri: os.Getenv("POSTGRES_URI"),
	}, nil
}

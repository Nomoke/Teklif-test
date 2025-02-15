package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Database Database
	Logger   Logger
}

type Database struct {
	PgUser     string `env:"POSTGRES_USER"`
	PgPassword string `env:"POSTGRES_PASSWORD"`
	PgHost     string `env:"POSTGRES_HOST"`
	PgPort     uint16 `env:"POSTGRES_PORT"`
	PgDatabase string `env:"POSTGRES_DB"`
	PgSSLMode  string `env:"POSTGRES_SSL_MODE"`
}

type Logger struct {
	LogFilePath string `env:"LOG_FILE_PATH"`
	Level       string `env:"LOG_LVL" envDefault:"debug"`
}

func New() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("failed to load .env file: %w", err)
	}

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse config from environment variables: %w", err)
	}

	return cfg, nil
}

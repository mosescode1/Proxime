package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

var AppConfig *Config

type Config struct {
	// Database connection string
	APP_NAME string `env:"APP_NAME" env-default:"Proxime"`
	APP_PORT string `env:"APP_PORT" env-default:"8080"`
	APP_ENV  string `env:"APP_ENV" env-default:"development"`
}

func LoadConfig(filePath string) (*Config, error) {

	AppConfig = &Config{}
	// Load config from file
	if err := cleanenv.ReadConfig(filePath, AppConfig); err != nil {
		slog.Error("Error loading config: ", "error", err)
		return nil, err
	}

	return AppConfig, nil
}

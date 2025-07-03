package configs

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log/slog"
)

type Config struct {
	ServerPort  string `env:"SERVER_PORT" envDefault:"8080"`
	LoggerLevel string `env:"LOGGER_LEVEL" envDefault:"info"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		slog.Error("error parsing env", "error", err)
	}
	return &cfg
}

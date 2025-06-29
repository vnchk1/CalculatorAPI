package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}
type LoggerConfig struct {
	Level string `yaml:"level"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config.yaml: %v", err)
	}
	return &cfg, nil
}

func ConvertLogLevel(lvlStr string) slog.Level {
	switch lvlStr {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

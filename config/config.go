package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		GRPC  `yaml:"grpc"`
		Log   `yaml:"logger"`
		Redis `yaml:"redis"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// GRPC -.
	GRPC struct {
		Port    string `env-required:"true" yaml:"port" env:"GRPC_PORT"`
		Timeout string `env-required:"true" yaml:"timeout" env:"GRPC_TIMEOUT"`
	}

	Redis struct {
		Host string `env-required:"true" yaml:"host" env:"REDIS_HOST"`
		Port string `env-required:"true" yaml:"port" env:"REDIS_PORT"`
	}
	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

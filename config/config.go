package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type ENV string

const (
	ENV_TEST ENV = "TEST"
	ENV_DEV  ENV = "DEV" // Fixed: Removed space in "DEV D"
)

type Config struct {
	DatabaseName     string `env:"DB_NAME"`
	DatabaseHost     string `env:"DB_HOST"`
	DatabaseUser     string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`
	DatabasePort     string `env:"DB_PORT"`
	DatabasePortTest string `env:"DB_PORT_TEST"`
	ProjectRoot      string `env:"PROJECT_ROOT"`
	Env              ENV    `env:"ENV" envDefault:"DEV"` // Ensure consistency
}

func (c *Config) DatabaseUrl() string {
	port := c.DatabasePort
	if c.Env == ENV_TEST {
		port = c.DatabasePortTest
	}

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		port,
		c.DatabaseName)
}

func New() (*Config, error) {
	cfg, err := env.ParseAs[Config]() // Ensure you have env v11+
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &cfg, nil
}

package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type ENV string

const (
	ENV_TEST ENV = "TEST"
	ENV_DEV  ENV = "DEV D"
)

type Config struct {
	DatabaseName     string `env:"DB_NAME"`
	DatabaseHost     string `env:"DB_HOST"`
	DatabaseUser     string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`
	DatabasePort     string `env:"DB_PORT"`
	DatabasePortTest string `env:"DB_PORT_TEST"`
	Env              ENV    `env:"ENV" envDefault: "DEV"`
}

func (c *Config) DatabaseUrl() string {

	PORT := c.DatabasePort

	if c.Env == ENV_TEST {
		PORT = c.DatabasePortTest
	}

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		PORT,
		c.DatabaseName)
}

func New() (*Config, error) {

	cfg, err := env.ParseAs[Config]()

	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}

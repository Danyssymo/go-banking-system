package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	HTTPPort    int    `env:"HTTP_PORT" envDefault:"8081"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

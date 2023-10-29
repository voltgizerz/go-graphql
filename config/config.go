package config

import (
	"github.com/go-graphql/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Database `yaml:"database"`
	}

	App struct {
		Name          string `env-required:"true" yaml:"name"`
		Version       string `env-required:"true" yaml:"version"`
		IsUseFakeData bool   `env-required:"true" yaml:"is_use_fake_data"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
	}

	Database struct {
		PoolMax int `env-required:"true" yaml:"pool_max"`
	}
)

// NewConfig returns app config.
func NewConfig() *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config-development.yml", cfg) // Change this base on your env
	if err != nil {
		logger.Log.Fatalln(err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	return cfg
}

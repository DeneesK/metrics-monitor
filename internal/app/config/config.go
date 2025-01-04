package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Address                 string        `env-default:"localhost:8000"`
	Timeout                 time.Duration `env-default:"4s"`
	LogEnv                  string        `env-default:"local"`
	ScrapingInterval        time.Duration `env-default:"4s"`
	GracefulShutdownTimeout time.Duration `env-default:"4s"`
}

func MustLoad() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}

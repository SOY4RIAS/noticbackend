package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Address string `env:"ADDRESS" envDefault:":3000"`
	JwtSecret string `env:"SEED_TOKEN,required"`
	DatabaseName string `env:"DATABASE_NAME,required"`
	DatabaseUri string `env:"DATABASE_URI,required"`
}

func New(files... string) *Config  {
	err := godotenv.Load(files...)

	if err != nil {
		log.Printf("environment File can't be found at %q\n", files)
	}

	config := Config{}


	err = env.Parse(&config)

	return &config
}
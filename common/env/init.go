package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port           int    `env:"PORT,unset" envDefault:"4000"`
	RootPass       string `env:"ROOT_PASSWORD,unset" envDefault:"root"`
	AccessTokenKey string `env:"ACCESS_TOKEN_KEY,unset"`
}

func LoadConfig() *config {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}

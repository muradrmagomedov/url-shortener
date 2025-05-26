package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Addr         string `env:"SERVER_ADDRESS"`
	ShortURLAddr string `env:"BASE_URL"`
}

func NewConfig() Config {
	ctf := Config{}
	env.Parse(&ctf)
	Addr := flag.String("a", "localhost:8080", "host address")
	ShortURLAddr := flag.String("b", "http://localhost:8080", "prefix for short url")
	flag.Parse()
	if ctf.Addr == "" {
		ctf.Addr = *Addr
	}
	if ctf.ShortURLAddr == "" {
		ctf.ShortURLAddr = *ShortURLAddr
	}
	return ctf
}

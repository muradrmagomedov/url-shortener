package config

import "flag"

type Config struct {
	Addr         string
	ShortURLAddr string
}

func NewConfig() Config {
	ctf := Config{}
	Addr := flag.String("a", "localhost:8080", "host address")
	ShortURLAddr := flag.String("b", "http://localhost:8080/", "prefix for short url")
	flag.Parse()
	ctf.Addr = *Addr
	ctf.ShortURLAddr = *ShortURLAddr
	return ctf
}

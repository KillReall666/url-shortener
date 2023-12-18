package config

import (
	"flag"
	"github.com/caarlos0/env/v10"
	"log"
)

type RunConfig struct {
	Address         string
	ShortURLAddress string
}

const (
	defaultAddress         = "" //":8080"
	defaultShortURLAddress = "" //"http://localhost:8080/"
)

func LoadConfig() RunConfig {
	cfg := RunConfig{}

	flag.StringVar(&cfg.Address, "a", defaultAddress, "server address")
	flag.StringVar(&cfg.ShortURLAddress, "b", defaultShortURLAddress, "returned link address")

	err := env.Parse(&cfg)
	if err != nil {
		log.Println(err)
	}

	return cfg
}

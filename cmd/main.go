package main

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
	"github.com/devkekops/ctf-signature/internal/app/config"
	"github.com/devkekops/ctf-signature/internal/app/server"
)

func main() {
	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "server address")
	flag.StringVar(&cfg.SecretKey, "k", cfg.SecretKey, "secret key")
	flag.StringVar(&cfg.Flag, "f", cfg.Flag, "flag")
	flag.Parse()

	log.Fatal(server.Serve(&cfg))
}

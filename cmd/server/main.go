package main

import (
	"flag"
	"log"

	_ "github.com/lib/pq"

	"github.com/opencars/koatuu/pkg/apiserver"
	"github.com/opencars/koatuu/pkg/config"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	// Get configuration.
	conf, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(conf); err != nil {
		log.Fatal(err)
	}
}

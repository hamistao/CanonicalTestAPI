package main

import (
	"canonicalTestAPI/pkg/config"
	"canonicalTestAPI/pkg/routes"
	"flag"
	"fmt"
	"log"
)

var flagConfig = flag.String("config", "./config.yml", "path to the config file")

func main() {
	cfg, err := config.Load(*flagConfig)
	if err != nil {
		log.Fatal(err)
	}

	router := routes.Router()

	router.Run(fmt.Sprintf("localhost:%d", cfg.ServerPort))
}

package main

import (
	"canonicalTestAPI/pkg/config"
	"canonicalTestAPI/pkg/routes"
	"canonicalTestAPI/pkg/service"
	"flag"
	"fmt"
	"log"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

var flagConfig = flag.String("config", "./config.yml", "path to the config file")

func main() {
	cfg, err := config.Load(*flagConfig)
	if err != nil {
		log.Fatal(err)
	}

	db, err := dbx.MustOpen("postgres", cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}
	sv := service.Service{
		DB: db,
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	router := routes.Router(sv)

	router.Run(fmt.Sprintf("localhost:%d", cfg.ServerPort))
}

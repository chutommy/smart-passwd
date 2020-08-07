package main

import (
	"log"

	config "github.com/chutified/smart-passwd/config"
	server "github.com/chutified/smart-passwd/server"
	_ "github.com/lib/pq"
)

func main() {

	// get configuration
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// set server
	srv := server.New()
	err = srv.Set(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// run server
	err = srv.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Panic(srv.Stop())
	}()
}

package main

import (
	"log"

	"github.com/chutified/smart-passwd/config"
	"github.com/chutified/smart-passwd/server"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// get configuration
	cfg, err := config.GetConfig("config.yml")
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

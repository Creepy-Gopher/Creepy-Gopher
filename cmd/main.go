package main

import (
	"creepy/service"
	"log"
)

func main() {
	//cfg := readConfig()

	app, err := service.NewAppContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//http_server.Run(cfg.Server, app)
	//go run http client crawller
	//run bot client
}

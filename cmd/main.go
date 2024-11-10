package main

import (
	"creepy/service"
	"log"
	"flag"
	"fmt"
	"creepy/config"
)

func main() {

	envPath := flag.String("env", "", "path of env file");
	flag.Parse()

	fmt.Println("env path:", *envPath)

	cfg := config.ReadConfig(*envPath)

	// app, err := service.NewAppContainer(cfg)
	_, err := service.NewAppContainer(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	//http_server.Run(cfg.Server, app)
	//go run http client crawller
	//run bot client
}

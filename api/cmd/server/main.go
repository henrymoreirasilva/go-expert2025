package main

import (
	"log"

	"github.com/henrymoreirasilva/go-api/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	println(config.DBDriver)
}

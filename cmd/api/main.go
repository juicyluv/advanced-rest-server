package main

import (
	"log"

	"github.com/juicyluv/advanced-rest-server/internal/server"
)

func main() {
	config := server.NewConfig()
	server := server.New(config)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

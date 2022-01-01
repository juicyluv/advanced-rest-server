package main

import (
	"flag"
	"log"

	"github.com/juicyluv/advanced-rest-server/config"
	"github.com/juicyluv/advanced-rest-server/internal/server"
)

var (
	configPath = flag.String("config-path", "config", "config directory path")
	configName = flag.String("config-name", "server", "config filename without extension")
)

func main() {
	flag.Parse()

	if err := config.LoadConfigs(*configPath, *configName); err != nil {
		log.Fatalln("An error occurred while loading config file. Please, check config path and filename.")
	}

	config := server.NewConfig()
	server := server.New(config)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"log"

	"github.com/juicyluv/advanced-rest-server/config"
	"github.com/juicyluv/advanced-rest-server/internal/logger"
	"github.com/juicyluv/advanced-rest-server/internal/server"

	_ "github.com/lib/pq"
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

	logger := logger.New(logger.LevelDebug)

	config := server.NewConfig()
	server := server.New(config, logger)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

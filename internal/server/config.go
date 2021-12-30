package server

import "github.com/juicyluv/advanced-rest-server/internal/logger"

type config struct {
	Port     string
	LogLevel int
}

func NewConfig() *config {
	return &config{
		Port:     "3000",
		LogLevel: logger.LevelDebug,
	}
}

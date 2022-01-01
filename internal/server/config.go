package server

import (
	"time"

	"github.com/juicyluv/advanced-rest-server/internal/logger"
	"github.com/spf13/viper"
)

type db struct {
	Username string
	DbName   string
	Port     string
	Host     string
	SSLMode  string
}

type config struct {
	Db             db
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	LogLevel       logger.LogLevel
	MaxHeaderBytes int
}

func NewConfig() *config {
	return &config{
		Db: db{
			Port:     viper.GetString("db.postgres.port"),
			Username: viper.GetString("db.postgres.username"),
			DbName:   viper.GetString("db.postgres.dbname"),
			Host:     viper.GetString("db.postgres.host"),
			SSLMode:  viper.GetString("db.postgres.sslmode"),
		},
		Port:     viper.GetString("http.port"),
		LogLevel: logger.LevelDebug,
	}
}

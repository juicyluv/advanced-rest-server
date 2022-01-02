package server

import (
	"fmt"
	"os"
	"time"

	"github.com/juicyluv/advanced-rest-server/internal/logger"
	"github.com/spf13/viper"
)

type db struct {
	Username        string
	Password        string
	DbName          string
	Port            string
	Host            string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxIdleTimeConn time.Duration
}

type Config struct {
	Db             db
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	LogLevel       logger.LogLevel
	MaxHeaderBytes int
}

// NewConfig return a new config instance.
func NewConfig() *Config {
	return &Config{
		Db: db{
			Password:        os.Getenv("POSTGRES_PASSWORD"),
			Username:        viper.GetString("db.postgres.username"),
			Port:            viper.GetString("db.postgres.port"),
			DbName:          viper.GetString("db.postgres.dbname"),
			Host:            viper.GetString("db.postgres.host"),
			SSLMode:         viper.GetString("db.postgres.sslmode"),
			MaxOpenConns:    viper.GetInt("db.postgres.maxOpenConns"),
			MaxIdleConns:    viper.GetInt("db.postgres.maxIdleConns"),
			MaxIdleTimeConn: time.Minute * time.Duration(viper.GetInt("db.postgres.maxIdleConns")),
		},
		Port:     viper.GetString("http.port"),
		LogLevel: logger.LevelDebug,
	}
}

// GetPostgresDSN returns a PostgreSQL DSN.
func (c *Config) GetPostgresDSN() string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Db.Username, c.Db.Password, c.Db.Host, c.Db.Port, c.Db.DbName, c.Db.SSLMode)
	return dsn
}

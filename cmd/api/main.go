package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/juicyluv/advanced-rest-server/configs"
	"github.com/juicyluv/advanced-rest-server/internal/logger"
	"github.com/juicyluv/advanced-rest-server/internal/server"

	_ "github.com/lib/pq"
)

var (
	configPath = flag.String("config-path", "configs", "config directory path")
	configName = flag.String("config-name", "server", "config filename without extension")
)

func main() {
	flag.Parse()

	if err := configs.LoadConfigs(*configPath, *configName); err != nil {
		log.Fatal(err)
	}

	logger := logger.New(logger.LevelDebug)
	config := server.NewConfig()

	db, err := openDB(config)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer db.Close()
	logger.Info("Connected to database '" + config.Db.DbName + "'.")

	server := server.New(config, logger, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// openDB opens a database connection. On success returns a pointer to DB connection.
func openDB(config *server.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", config.GetPostgresDSN())
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open (in-use + idle) connections in the pool.
	db.SetMaxOpenConns(config.Db.MaxOpenConns)

	// Set the maximum number of idle connections in the pool.
	db.SetMaxIdleConns(config.Db.MaxIdleConns)

	// Set the maximum idle connection timeout.
	db.SetConnMaxIdleTime(config.Db.MaxIdleTimeConn)

	// Create context for waiting to DB conenction
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// If DB connection cannot be estabilished within 5 seconds, return an error
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

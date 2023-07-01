package cmd

import (
	"flag"
	"fmt"
	"log"

	customlogger "travelo/internal/custom_logger"
	"travelo/internal/database"
	"travelo/internal/env"
	"travelo/internal/graphql"
	"travelo/internal/version"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
)

func Run(logger *log.Logger) error {
	var cfg Config

	cfg.BaseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.HttpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.DB.DSN = env.GetString("DB_DSN", "daffashafwan:daffashafwan@localhost:5432/dbtravelo?sslmode=disable")
	cfg.DB.AutoMigrate = env.GetBool("DB_AUTOMIGRATE", true)

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.DB.DSN, cfg.DB.AutoMigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	validate := validator.New()

	gql := graphql.NewGraphqlClient("https://exciting-deer-66.hasura.app/v1/graphql")

	customLog, err := customlogger.NewLogger("app.log")
	if err != nil {
		return err
	}

	client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
    })

	app := &Application{
		Config: cfg,
		DB:     db,
		Logger: logger,
		Validator: validate,
		GraphqlClient: gql,
		CustomLogger: customLog,
		Redis: client,
	}

	return app.serveHTTP()
}
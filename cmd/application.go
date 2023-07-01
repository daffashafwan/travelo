package cmd

import (
	"log"
	"sync"

	"travelo/internal/custom_logger"
	"travelo/internal/database"
	"travelo/internal/graphql"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
)

type Config struct {
	BaseURL  string
	HttpPort int
	DB       struct {
		DSN         string
		AutoMigrate bool
	}
}

type Application struct {
	Config        Config
	DB            *database.DB
	Logger        *log.Logger
	wg            sync.WaitGroup
	Validator     *validator.Validate
	GraphqlClient *graphql.GraphqlClient
	CustomLogger  *customlogger.Logger
	Redis         *redis.Client
}

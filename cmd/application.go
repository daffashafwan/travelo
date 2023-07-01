package cmd

import (
	"log"
	"sync"

	"travelo/internal/database"
	"github.com/go-playground/validator/v10"
	"travelo/internal/graphql"
	"travelo/internal/custom_logger"
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
	Config    Config
	DB        *database.DB
	Logger    *log.Logger
	wg        sync.WaitGroup
	Validator *validator.Validate
	GraphqlClient *graphql.GraphqlClient
	CustomLogger *customlogger.Logger
}
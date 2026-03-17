package main

import (
	"log"

	"github.com/DrxwDev/rest-api/internal/config"
	"github.com/DrxwDev/rest-api/internal/database"
	"github.com/DrxwDev/rest-api/internal/logger"
	"github.com/DrxwDev/rest-api/internal/server"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	fx.New(
		config.Module,
		database.Module,
		server.Module,
		logger.Module,
	).Run()
}

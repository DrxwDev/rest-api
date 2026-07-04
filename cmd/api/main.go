package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/DrxwDev/rest-api/internal/config"
	"github.com/DrxwDev/rest-api/internal/database"
	"github.com/DrxwDev/rest-api/internal/logger"
	"github.com/DrxwDev/rest-api/internal/server"
)

func main() {
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found")
		}
	}

	fx.New(
		config.Module,
		database.Module,
		server.Module,
		logger.Module,
	).Run()
}

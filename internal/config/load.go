package config

import (
	"os"
	"time"
)

func LoadApp() AppConfig {
	return AppConfig{
		HOST: os.Getenv("HOST"),
		PORT: os.Getenv("PORT"),
	}
}

func LoadDB() DBConfig {
	return DBConfig{
		URL: os.Getenv("DB_URL"),
	}
}

func LoadServer() ServerConfig {
	return ServerConfig{
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    2 << 10,
	}
}

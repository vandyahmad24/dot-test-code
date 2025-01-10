package config

import (
	"log"

	"github.com/vandyahmad24/alat-bantu/config"
)

var (
	configPath = "./.env"
	EnvConfig  = NewConfig()
)

type Config struct {
	PORT string

	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_DATABASE string

	REDIS_HOST string
	REDIS_PASS string
	JWT_SECRET string
}

func NewConfig() *Config {
	if err := config.LoadEnv(configPath); err != nil {
		log.Fatalf("Application dimissed. Application cannot find %s to run this application", err.Error())
	}

	return &Config{
		PORT: config.GetEnv("PORT", ""),

		MYSQL_USERNAME: config.GetEnv("MYSQL_USERNAME", ""),
		MYSQL_PASSWORD: config.GetEnv("MYSQL_PASSWORD", ""),
		MYSQL_HOST:     config.GetEnv("MYSQL_HOST", ""),
		MYSQL_PORT:     config.GetEnv("MYSQL_PORT", ""),
		MYSQL_DATABASE: config.GetEnv("MYSQL_DATABASE", ""),

		JWT_SECRET: config.GetEnv("JWT_SECRET", ""),

		REDIS_HOST: config.GetEnv("REDIS_HOST", ""),
		REDIS_PASS: config.GetEnv("REDIS_PASS", ""),
	}

}

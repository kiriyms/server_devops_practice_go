package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Address     string
	Port        string
	Environment string
}

func MustLoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	address, ok := os.LookupEnv("ADDRESS")
	if !ok {
		log.Fatal("ADDRESS environment variable missing")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT environment variable missing")
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		log.Fatal("ENVIRONMENT environment variable missing")
	}

	return &Config{
		Address:     address,
		Port:        port,
		Environment: env,
	}
}

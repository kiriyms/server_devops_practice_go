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

	var ok bool
	var address string
	if address, ok = os.LookupEnv("ADDRESS"); !ok {
		log.Fatal("ADDRESS environment varriable missing")
	}
	var port string
	if port, ok = os.LookupEnv("PORT"); !ok {
		log.Fatal("PORT environment varriable missing")
	}
	var env string
	if env, ok = os.LookupEnv("ENVIRONMENT"); !ok {
		log.Fatal("ENVIRONMENT environment varriable missing")
	}

	return &Config{
		Address:     address,
		Port:        port,
		Environment: env,
	}
}

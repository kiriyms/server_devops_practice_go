package common

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Address     string
	Port        string
	Environment string
}

var (
	cfg     *Config
	onceCfg sync.Once
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

func MustLoadConfig() {
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

	onceCfg.Do(func() {
		cfg = &Config{
			Address:     address,
			Port:        port,
			Environment: env,
		}
	})
}

func GetConfig() *Config {
	if cfg == nil {
		panic("Global config not initialized. Call MustLoadConfig() first.")
	}
	return cfg
}

package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        int
	DatabaseURL string
}

func Load(files ...string) (*Config, error) {
	var err error
	if files == nil {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(files[0])
	}

	if err != nil {
		return nil, err
	}

	var config Config

	config.Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	switch os.Getenv("API_MODE") {
	case "dev":
		config.DatabaseURL = os.Getenv("DEV_DB_URL")
	case "test":
		config.DatabaseURL = os.Getenv("TEST_DB_URL")
	case "prod":
		config.DatabaseURL = os.Getenv("PROD_DB_URL")
	}

	return &config, nil
}

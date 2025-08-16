package config

import (
	"context"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Config struct {
	Port int
	DB   *pgxpool.Pool
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

	var connString string
	switch os.Getenv("API_MODE") {
	case "dev":
		connString = os.Getenv("DEV_DB_URL")
	case "test":
		connString = os.Getenv("TEST_DB_URL")
	case "prod":
		connString = os.Getenv("PROD_DB_URL")
	}

	if config.DB, err = pgxpool.New(context.Background(), connString); err != nil {
		return nil, err
	}

	return &config, nil
}

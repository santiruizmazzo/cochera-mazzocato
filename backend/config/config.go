package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT int
}

func Load(files ...string) (*Config, error) {
	err := godotenv.Load(files[0])
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	var config Config
	config.PORT = port
	return &config, nil
}

package config

import (
	"log"
	"os"
	"testing"
)

var config *Config

const ENV_FILE_NAME = ".env.test"
const ENV_FILE_CONTENT = "PORT=5000"

func TestMain(m *testing.M) {
	err := os.WriteFile(ENV_FILE_NAME, []byte(ENV_FILE_CONTENT), 0666)
	if err != nil {
		log.Fatal(err)
	}

	config, err = Load(ENV_FILE_NAME)
	if err != nil {
		log.Fatalf("%s: %s", err, "No se pudo cargar ninguna variable")
	}

	code := m.Run()

	err = os.Remove(ENV_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func TestLoadPort(t *testing.T) {
	if config.PORT != 5000 {
		t.Errorf("Puerto cargado %v != 5000", config.PORT)
	}
}

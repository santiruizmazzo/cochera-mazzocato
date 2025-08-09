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

	err = os.Remove(ENV_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	os.Clearenv()

	code := m.Run()

	os.Exit(code)
}

func TestDefaultEnvFileWhenNonePassed(t *testing.T) {
	err := os.WriteFile(".env", []byte("PORT=777"), 0666)
	if err != nil {
		t.Errorf("Error creando archivo temporal .env: %s", err)
	}

	defaultConfig, _ := Load()
	os.Remove(".env")

	if defaultConfig.Port != 777 {
		t.Errorf("Puerto cargado: %v != 777", defaultConfig.Port)
	}
}

func TestLoadedPortIsCorrect(t *testing.T) {
	if config.Port != 5000 {
		t.Errorf("Puerto cargado %v != 5000", config.Port)
	}
}

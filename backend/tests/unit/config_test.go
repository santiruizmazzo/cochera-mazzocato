package unit

import (
	"cochera/application/config"
	"log"
	"os"
	"testing"
)

const envFileContent = `PORT=777`

func TestMain(m *testing.M) {
	err := os.WriteFile(".env", []byte(envFileContent), 0666)
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	err = os.Remove(".env")
	if err != nil {
		log.Fatalf("Failed deleting temp file .env: %s", err)
	}
	os.Exit(code)
}

func TestDefaultEnvFileWhenNonePassed(t *testing.T) {
	defaultConfig, _ := config.Load()
	if defaultConfig.Port != 777 {
		t.Fatalf("Loaded port %v != 777", defaultConfig.Port)
	}

	os.Clearenv()
}

func TestLoadWithCustomEnvFile(t *testing.T) {
	err := os.WriteFile(".env.test", []byte("PORT=666"), 0666)
	if err != nil {
		t.Fatalf("Failed creating temp file .env.test: %s", err)
	}

	config, _ := config.Load(".env.test")
	if config.Port != 666 {
		t.Fatalf("Loaded port %v != 666", config.Port)
	}

	err = os.Remove(".env.test")
	if err != nil {
		t.Fatalf("Failed deleting temp file .env: %s", err)
	}
	os.Clearenv()
}

func TestLoadedPortIsCorrect(t *testing.T) {
	err := os.Setenv("PORT", "1234")
	if err != nil {
		t.Fatalf("Failed setting PORT as environment variable: %s", err)
	}

	config, err := config.Load()
	if err != nil {
		t.Fatalf("Failed loading configuration: %s", err)
	}

	if config.Port != 1234 {
		t.Fatalf("Loaded port %v != 1234", config.Port)
	}

	os.Clearenv()
}

func TestDatabaseURLWhenAPIModeIsDev(t *testing.T) {
	err := os.Setenv("API_MODE", "dev")
	if err != nil {
		t.Fatalf("Failed setting API_MODE as environment variable: %s", err)
	}

	const expectedDatabaseURL = "postgres://postgres:example@dev_db:5433/postgres"
	err = os.Setenv("DEV_DB_URL", expectedDatabaseURL)
	if err != nil {
		t.Fatalf("Failed setting DEV_DB_URL as environment variable: %s", err)
	}

	config, _ := config.Load()
	resultDatabaseURL := config.DB.Config().ConnString()
	if resultDatabaseURL != expectedDatabaseURL {
		t.Fatalf("Loaded URL %s != %s", resultDatabaseURL, expectedDatabaseURL)
	}

	os.Clearenv()
}

func TestDatabaseURLWhenAPIModeIsTest(t *testing.T) {
	err := os.Setenv("API_MODE", "test")
	if err != nil {
		t.Fatalf("Failed setting API_MODE as environment variable: %s", err)
	}

	const expectedDatabaseURL = "postgres://postgres:example@test_db:5434/postgres"
	err = os.Setenv("TEST_DB_URL", expectedDatabaseURL)
	if err != nil {
		t.Fatalf("Failed setting TEST_DB_URL as environment variable: %s", err)
	}

	config, _ := config.Load()
	resultDatabaseURL := config.DB.Config().ConnString()
	if resultDatabaseURL != expectedDatabaseURL {
		t.Fatalf("Loaded URL %s != %s", resultDatabaseURL, expectedDatabaseURL)
	}

	os.Clearenv()
}

func TestDatabaseURLWhenAPIModeIsProd(t *testing.T) {
	err := os.Setenv("API_MODE", "prod")
	if err != nil {
		t.Fatalf("Failed setting API_MODE as environment variable: %s", err)
	}

	const expectedDatabaseURL = "postgres://postgres:example@prod_db:5435/postgres"
	err = os.Setenv("PROD_DB_URL", expectedDatabaseURL)
	if err != nil {
		t.Fatalf("Failed setting PROD_DB_URL as environment variable: %s", err)
	}

	config, _ := config.Load()
	resultDatabaseURL := config.DB.Config().ConnString()
	if resultDatabaseURL != expectedDatabaseURL {
		t.Fatalf("Loaded URL %s != %s", resultDatabaseURL, expectedDatabaseURL)
	}

	os.Clearenv()
}

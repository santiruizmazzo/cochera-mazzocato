package postgres

import (
	"cochera/internal/domain/tenant"
	"cochera/internal/domain/time"
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool
var err error

func setupTestDatabase() error {
	db, err = pgxpool.New(context.Background(), os.Getenv("TEST_DB_URL"))
	if err != nil {
		return err
	}

	return nil
}

func cleanupTestDatabase() {
	_, err = db.Exec(context.Background(), `
		DO
		$func$
		BEGIN
			EXECUTE (
				SELECT 'TRUNCATE TABLE ' || string_agg(format('%I.%I', schemaname, tablename), ', ')
					|| ' RESTART IDENTITY CASCADE'
				FROM pg_tables
				WHERE schemaname = 'public'
			);
		END
		$func$;
	`)
}

func cleanupAndCloseTestDatabase() {
	cleanupTestDatabase()
	db.Close()
}

func TestMain(m *testing.M) {
	code := 1

	defer func() {
		os.Exit(code)
	}()

	err = setupTestDatabase()
	if err != nil {
		log.Printf("Error al conectar con DB de tests: %v", err)
		return
	}
	defer cleanupAndCloseTestDatabase()

	code = m.Run()
}

func TestPostgresTenantRepository_Save_Successfully_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repo := NewPostgresTenantRepository(db)

	localTenant := &tenant.Tenant{
		DNI:        12345678,
		Name:       "Manolo",
		LastName:   "Lamas",
		Address:    "Avenida Siempreviva 555",
		Phone:      "+5645551114",
		Email:      "mlamas@fifa09.com",
		EntryMonth: time.NewMonthOfYear(8, 2025),
	}

	savedTenant, err := repo.Save(localTenant)
	if err != nil {
		t.Fatal(err)
	}

	localTenant.ID = 1
	if savedTenant != localTenant {
		t.Fatal("El inquilino en memoria y el almacenado son diferentes")
	}
}

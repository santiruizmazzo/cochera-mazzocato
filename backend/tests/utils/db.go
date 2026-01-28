package utils

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupTestDatabase() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), os.Getenv("TEST_DB_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CleanupTestDatabase(db *pgxpool.Pool) {
	ctx := context.Background()

	// 1. Liberar slots
	_, err := db.Exec(ctx, `UPDATE slots SET tenant_id = NULL`)
	if err != nil {
		log.Println("⚠️ Failed to clear slots: ", err)
		return
	}

	// 2. DELETE en lugar de TRUNCATE (evita problemas de FK)
	_, err = db.Exec(ctx, `DELETE FROM tenants`)
	if err != nil {
		log.Println("⚠️ Failed to delete tenants: ", err)
		return
	}

	// 3. Reiniciar secuencia
	_, err = db.Exec(ctx, `ALTER SEQUENCE tenants_id_seq RESTART WITH 1`)
	if err != nil {
		log.Println("⚠️ Failed to reset sequence: ", err)
		return
	}

	// 4. Verificar slots
	var slotCount int
	err = db.QueryRow(ctx, `SELECT COUNT(*) FROM slots`).Scan(&slotCount)

	if err != nil {
		log.Println("⚠️ Failed to scan slot count from DB: ", err)
		return
	}

	if slotCount != 12 {
		log.Printf("⚠️ Only %d slots, re-seeding...", slotCount)
		return
	}
}

func CleanupAndCloseTestDatabase(db *pgxpool.Pool) {
	CleanupTestDatabase(db)
	db.Close()
}

func ClearTenantsTable(db *pgxpool.Pool) {
	ctx := context.Background()

	// Liberar slots
	_, err := db.Exec(ctx, `UPDATE slots SET tenant_id = NULL`)
	if err != nil {
		log.Println("failed to clear slots: ", err)
		return
	}

	// DELETE tenants (respeta FK constraints)
	_, err = db.Exec(ctx, `DELETE FROM tenants`)
	if err != nil {
		log.Println("failed to delete tenants: ", err)
		return
	}

	// Reiniciar secuencia
	_, err = db.Exec(ctx, `ALTER SEQUENCE tenants_id_seq RESTART WITH 1`)
	if err != nil {
		log.Println("failed to reset sequence: ", err)
		return
	}
}

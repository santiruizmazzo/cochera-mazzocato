package utils

import (
	"context"
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
	_, _ = db.Exec(context.Background(), `
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

func CleanupAndCloseTestDatabase(db *pgxpool.Pool) {
	CleanupTestDatabase(db)
	db.Close()
}

package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:5432/?sslmode=disable"
)

func ConnectDB() (*sql.DB, error) {
	// Crear base en disco (persistente)
	conn, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	var exists int
	err = conn.QueryRow(`SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='products';`).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		if _, err := os.Stat("db/schema/schema.sql"); err == nil {
			schema, _ := os.ReadFile("db/schema/schema.sql")
			_, err = conn.Exec(string(schema))
			if err != nil {
				return nil, err
			}
		}
	}

	return conn, conn.Ping()
}

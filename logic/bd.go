package logic

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(inTest bool) *sql.DB {
	if inTest {
		return connectTestDB()
	} else {
		return connectNormalDB()
	}
}

func connectNormalDB() *sql.DB {
	connStr := "user=videla password='XYZ' dbname=tarjetasdb port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("No se pudo establecer conexión con la base de datos: %v", err)
	}
	return db
}

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:5432/?sslmode=disable"
)

func connectTestDB() *sql.DB {
	// Crear base en disco (persistente)
	conn, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos de PRUEBA: %v", err)
	}

	var exists int
	err = conn.QueryRow(`SELECT COUNT(*) FROM sqlite_master WHERE type='table' ;`).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	if exists == 0 {
		if _, err := os.Stat("db/schema/schema.sql"); err == nil {
			schema, _ := os.ReadFile("db/schema/schema.sql")
			_, err = conn.Exec(string(schema))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	err = conn.Ping()
	if err != nil {
		log.Fatalf("No se pudo establecer conexión con la base de datos: %v", err)
	}
	return conn
}

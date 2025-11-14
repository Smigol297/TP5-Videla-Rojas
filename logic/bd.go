package logic

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
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

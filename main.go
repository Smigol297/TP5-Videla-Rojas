package main

import (
	"fmt"
	"log"
	"net/http"
	"tp5ne/db"
	sqlc "tp5ne/db/sqlc"
	"tp5ne/handlers"
)

func initServer() {
	// Define el puerto y muestra un mensaje en consola
	port := ":8081"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	// Inicia el servidor HTTP
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err) // Muestra error si falla el inicio
	}
}

func main() {
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries := sqlc.New(conn)
	productHandler := handlers.NewProductHandler(queries)

	http.Handle("/", productHandler)

	initServer()
}

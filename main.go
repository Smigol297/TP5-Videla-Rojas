package main

import (
	"fmt"
	"net/http"
	sqlc "tp5/db"
	"tp5/handlers"
	"tp5/logic"
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
	var inTest = false
	conn := logic.ConnectDB(inTest)
	defer conn.Close()

	queries := sqlc.New(conn)
	tarjetaHandler := handlers.NewTarjetaHandler(queries)

	http.Handle("/", tarjetaHandler)
	http.HandleFunc("/temas", logic.TemasHandler)

	initServer()
}

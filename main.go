package main

import (
	"log"
	"net/http"
	sqlc "tp5/db"
	"tp5/handlers"
	"tp5/logic"
	"tp5/middleware"
	"tp5/views"
)

func main() {
	conn := logic.ConnectDB()
	defer conn.Close()

	queries := sqlc.New(conn)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		views.WelcomePage("Bienvenido").Render(r.Context(), w)
	})

	tar := handlers.NewTarjetaHandler(queries)
	tarHandler := middleware.LoggingMiddleware(middleware.AuthMiddleware(tar))
	mux.Handle("/tarjetas", tarHandler)
	mux.Handle("/tarjetas/", tarHandler)
	tema := handlers.NewTemaHandler(queries)
	temaHandler := middleware.LoggingMiddleware(middleware.AuthMiddleware(tema))
	mux.Handle("/temas", temaHandler)
	mux.Handle("/temas/", temaHandler)
	usuario := handlers.NewUsuarioHandler(queries)
	usuarioHandler := middleware.LoggingMiddleware(middleware.AuthMiddleware(usuario))
	mux.Handle("/usuarios", usuarioHandler)
	mux.Handle("/usuarios/", usuarioHandler)

	log.Println("starting server :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}

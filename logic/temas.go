package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	sqlc "tp5/db"

	_ "github.com/lib/pq"
)

// INICIO CON MINUSCULA = NO PUBLICO
// INICIO CON MAYUSCULA = PUBLICO
func ValidateCreateTema(nombre string) error {
	if nombre == "" {
		return fmt.Errorf("el nombre del tema no puede estar vacío")
	}
	return nil
}
func ValidateUpdateTema(p sqlc.UpdateTemaParams) error {
	if p.IDTema <= 0 {
		return fmt.Errorf("ID de tema %d inválido", p.IDTema)
	}
	if p.NombreTema == "" {
		return fmt.Errorf("el nombre del tema no puede estar vacío")
	}

	return nil
}

func getTemas(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	temas, err := queries.ListTemas(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Enviar los datos como JSON válido
	if err := json.NewEncoder(w).Encode(temas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func TemasHandler(w http.ResponseWriter, r *http.Request) {
	//esto se deberia sacar
	var inTest = false

	db := ConnectDB(inTest)
	defer db.Close()

	defer db.Close()
	queries := sqlc.New(db)
	switch r.Method {
	case http.MethodGet:
		getTemas(w, r, queries)
	case http.MethodPost:
		createTema(w, r, queries)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
func createTema(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries) {
	type reqBody struct {
		NombreTema string `json:"nombre_tema"`
	}

	var body reqBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// validar nombre con tu función existente
	if err := ValidateCreateTema(body.NombreTema); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	newTema, err := queries.CreateTema(ctx, body.NombreTema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newTema); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TemasByIDHandler(w http.ResponseWriter, r *http.Request) {
	//esto se deberia sacar
	var inTest = false

	db := ConnectDB(inTest)
	defer db.Close()

	defer db.Close()
	queries := sqlc.New(db)
	// Extraer el ID del tema de la URL
	var id int
	_, err := fmt.Sscanf(r.URL.Path, "/temas/%d", &id)
	if err != nil {
		http.Error(w, "ID de tema inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTemaByID(w, r, id, queries)
	case http.MethodPut:
		putTemaByID(w, r, id, queries)
	case http.MethodDelete:
		deleteTemaByID(w, r, id, queries)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
func getTemaByID(w http.ResponseWriter, r *http.Request, id int, queries *sqlc.Queries) {
	ctx := context.Background()
	tema, err := queries.GetTemaById(ctx, int32(id))
	if err != nil {
		http.Error(w, "Tema no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Enviar los datos como JSON válido
	if err := json.NewEncoder(w).Encode(tema); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func putTemaByID(w http.ResponseWriter, r *http.Request, id int, queries *sqlc.Queries) {
	var p sqlc.UpdateTemaParams
	p.IDTema = int32(id)
	// decodificar JSON
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	// validar tema
	if err := ValidateUpdateTema(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// actualizar tema en la base de datos
	ctx := context.Background()
	err = queries.UpdateTema(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Enviar los datos como JSON válido
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func deleteTemaByID(w http.ResponseWriter, r *http.Request, id int, queries *sqlc.Queries) {
	ctx := context.Background()
	err := queries.DeleteTema(ctx, int32(id))
	if err != nil {
		http.Error(w, "Tema no encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

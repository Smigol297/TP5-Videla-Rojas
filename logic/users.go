package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	sqlc "tp5/db"

	_ "github.com/lib/pq"
)

// INICIO CON MINUSCULA = NO PUBLICO
// INICIO CON MAYUSCULA = PUBLICO
func ValidateCreateUser(p sqlc.CreateUsuarioParams) error {
	if p.NombreUsuario == "" {
		return fmt.Errorf("el nombre del usuario no puede estar vacío")
	}
	_, err := mail.ParseAddress(p.Email)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("el email del usuario no es válido")
	}
	if p.Contrasena == "" {
		return fmt.Errorf("la contraseña del usuario no puede estar vacía")
	}
	return nil
}
func ValidateUpdateUser(p sqlc.UpdateUsuarioParams) error {
	if p.IDUsuario <= 0 {
		return fmt.Errorf("ID de usuario %d inválido", p.IDUsuario)
	}
	if p.NombreUsuario == "" {
		return fmt.Errorf("el nombre del usuario no puede estar vacío")
	}
	_, err := mail.ParseAddress(p.Email)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("el email del usuario no es válido")
	}
	if p.Contrasena == "" {
		return fmt.Errorf("la contraseña del usuario no puede estar vacía")
	}
	return nil
}

func getUsers(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	usuarios, err := queries.ListUsuarios(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Enviar los datos como JSON válido
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	queries := sqlc.New(db)
	// Revisa el MÉTODO HTTP (GET, POST, etc.)
	switch r.Method {
	// SI ES GET /users -> Listar todos los usuarios
	case http.MethodGet:
		getUsers(w, r, queries)
	// SI ES POST /users -> Crear un nuevo usuario
	case http.MethodPost:
		createUser(w, r, queries)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
func createUser(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries) {
	var p sqlc.CreateUsuarioParams

	// decodificar JSON
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// validar usuario
	if err := ValidateCreateUser(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// crear usuario en la base de datos
	ctx := context.Background()
	newUsuario, err := queries.CreateUsuario(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Enviar los datos como JSON válido
	if err := json.NewEncoder(w).Encode(newUsuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UsersByIDHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	defer db.Close()
	queries := sqlc.New(db)
	// Extraer el ID del usuario de la URL
	var id int
	_, err := fmt.Sscanf(r.URL.Path, "/users/%d", &id)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUserByID(w, r, id, queries)
	case http.MethodPut:
		putUserByID(w, r, id, queries)
	case http.MethodDelete:
		deleteUserByID(w, r, id, queries)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
func getUserByID(w http.ResponseWriter, r *http.Request, id int, queries *sqlc.Queries) {
	ctx := context.Background()
	user, err := queries.GetUsuarioById(ctx, int32(id))
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Enviar los datos como JSON válido
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func putUserByID(w http.ResponseWriter, r *http.Request, id int, queries *sqlc.Queries) {
	var p sqlc.UpdateUsuarioParams
	p.IDUsuario = int32(id)
	// decodificar JSON
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	// validar usuario
	if err := ValidateUpdateUser(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// actualizar usuario en la base de datos
	ctx := context.Background()
	err = queries.UpdateUsuario(ctx, p)
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
func deleteUserByID(w http.ResponseWriter, r *http.Request, id int, queries *sqlc.Queries) {
	ctx := context.Background()
	err := queries.DeleteUsuario(ctx, int32(id))
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

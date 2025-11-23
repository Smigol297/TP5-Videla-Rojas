package handlers

import (
	"context"
	"net/http"
	"tp5/db"
	"tp5/views"

	"fmt"
	"strconv"
	"strings"
	"tp5/logic"
)

type UsuarioHandler struct {
	queries *db.Queries
}

func NewUsuarioHandler(q *db.Queries) *UsuarioHandler {
	return &UsuarioHandler{queries: q}
}

func (h *UsuarioHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/usuarios") {
		http.NotFound(w, r)
		return
	}
	switch {
	case r.URL.Path == "/usuarios":
		switch r.Method {
		case http.MethodGet:
			h.GetUsuarios("Usuarios", w, r)
		//POST /usuarios
		case http.MethodPost:
			h.CreateUsuario(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	//parte usuario by id
	case strings.HasPrefix(r.URL.Path, "/usuarios/"):
		var id int
		_, err := fmt.Sscanf(r.URL.Path, "/usuarios/%d", &id)

		if err != nil {
			if idStr := r.URL.Query().Get("id"); idStr != "" {
				id, err = strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, "ID inválido", http.StatusBadRequest)
					return
				}
			}
		}
		switch r.Method {
		//GET/usuarios=1
		case http.MethodGet:
			h.GetUsuarioByID("Usuario por id", w, r, id)
		//PUT/usuarios=1
		case http.MethodPut:
			h.PutUsuarioByID(w, r, id)
		//DELETE/usuarios=1
		case http.MethodDelete:
			h.DeleteUsuarioByID(w, r, id)
		//para formularios que no soportan PUT o DELETE
		case http.MethodPost:
			switch r.FormValue("_method") {
			case "PUT":
				h.PutUsuarioByID(w, r, id)
			case "DELETE":
				h.DeleteUsuarioByID(w, r, id)
			default:
				http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			}
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	default:
		http.NotFound(w, r)
	}
}
func (h *UsuarioHandler) GetUsuarios(title string, w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.queries.ListUsuarios(context.Background())
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, usuarios, nil).Render(r.Context(), w)
}
func (h *UsuarioHandler) CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var p db.CreateUsuarioParams
	r.ParseForm()
	p.NombreUsuario = r.Form.Get("usuario-nombre")
	p.Email = r.Form.Get("usuario-email")
	p.Contrasena = r.Form.Get("usuario-contrasena")
	// validar usuario
	if err := logic.ValidateCreateUser(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	_, err := h.queries.CreateUsuario(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//http.Redirect(w, r, "/usuarios", http.StatusSeeOther)

	// 2. En lugar de Redirect, consultamos la lista actualizada
	usuarios, err := h.queries.ListUsuarios(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//  Renderizamos SOLO el componente de la lista (UserList)
	// HTMX tomará este HTML y reemplazará la tabla vieja en el navegador.
	views.UserList(usuarios).Render(r.Context(), w)
}

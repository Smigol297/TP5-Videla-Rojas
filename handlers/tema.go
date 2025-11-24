package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tp5/db"
	"tp5/logic"
	"tp5/views"
)

type TemaHandler struct {
	queries *db.Queries
}

func NewTemaHandler(q *db.Queries) *TemaHandler {
	return &TemaHandler{queries: q}
}

func (h *TemaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/temas") {
		http.NotFound(w, r)
		return
	}
	switch {
	case r.URL.Path == "/temas":
		switch r.Method {
		case http.MethodGet:
			h.GetTemas("Temas", w, r)
		//POST /temas
		case http.MethodPost:
			h.CreateTema(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	//parte tema by id
	case strings.HasPrefix(r.URL.Path, "/temas/"):
		var id int
		_, err := fmt.Sscanf(r.URL.Path, "/temas/%d", &id)
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
		//GET/temas=1
		case http.MethodGet:
			h.GetTemaByID("Tema por id", w, r, id)
		//PUT/temas=1
		case http.MethodPut:
			h.PutTemaByID(w, r, id)
		//DELETE/temas=1
		case http.MethodDelete:
			h.DeleteTemaByID(w, r, id)
		case http.MethodPost:
			switch r.FormValue("_method") {
			case "PUT":
				h.PutTemaByID(w, r, id)
			case "DELETE":
				h.DeleteTemaByID(w, r, id)
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

func (h *TemaHandler) GetTemas(title string, w http.ResponseWriter, r *http.Request) {
	tema, err := h.queries.ListTemas(context.Background())
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, tema, nil).Render(r.Context(), w)
}

func (h *TemaHandler) CreateTema(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var NombreTema = r.Form.Get("tema-nombre")

	// validar tarjeta
	if err := logic.ValidateCreateTema(NombreTema); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	_, err := h.queries.CreateTema(ctx, NombreTema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//http.Redirect(w, r, "/tarjetas", http.StatusSeeOther)

	//Obtener la lista actualizada de temas
	temas, err := h.queries.ListTemas(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Renderizar SOLO la lista de temas
	views.TemaList(temas).Render(r.Context(), w)
}

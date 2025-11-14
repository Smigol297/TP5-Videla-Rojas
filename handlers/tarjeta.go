package handlers

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	db "tp5/db"
	"tp5/logic"
	"tp5/views"
)

type TarjetaHandler struct {
	queries *db.Queries
}

func NewTarjetaHandler(q *db.Queries) *TarjetaHandler {
	return &TarjetaHandler{queries: q}
}
func (h *TarjetaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/tarjetas") {
		http.NotFound(w, r)
		return
	}
	switch r.URL.Path {
	case "/tarjetas":
		switch r.Method {
		case http.MethodGet:
			temaStr := r.URL.Query().Get("tema") // obtiene el parámetro "tema"
			//GET /tarjetas
			if temaStr == "" {
				h.GetTarjetasAndTemas("Lista de tarjetas", w, r)
				return
			}

			tema, err := strconv.Atoi(temaStr)
			if err != nil {
				http.Error(w, "ID de tarjeta inválido", http.StatusBadRequest)
				return
			}
			//GET /tarjetas?tema=1
			h.GetTarjetasByTema("Tarjetas por tema", tema, w, r)
		//POST /tarjetas
		case http.MethodPost:
			h.CreateTarjeta(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	//parte tarjetas by id
	case "/tarjetas/":
	default:
		http.NotFound(w, r)
	}
}

func (h *TarjetaHandler) GetTarjetas(title string, w http.ResponseWriter, r *http.Request) {
	tarjetas, err := h.queries.ListTarjetas(context.Background())
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, tarjetas, nil).Render(r.Context(), w)
}

func (h *TarjetaHandler) CreateTarjeta(w http.ResponseWriter, r *http.Request) {
	var p db.CreateTarjetaParams
	r.ParseForm()

	idTemaStr := r.Form.Get("id-tema")
	idTema, err := strconv.ParseInt(idTemaStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de tema inválido", http.StatusBadRequest)
		return
	}
	p.Pregunta = r.Form.Get("pregunta")
	p.Respuesta = r.Form.Get("respuesta")
	p.OpcionA = r.Form.Get("opcion-a")
	p.OpcionB = r.Form.Get("opcion-b")
	p.OpcionC = r.Form.Get("opcion-c")
	p.IDTema = int32(idTema)
	// validar tarjeta
	if err := logic.ValidateCreateTarjeta(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	_, err = h.queries.CreateTarjeta(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/tarjetas", http.StatusSeeOther)
}

func (h *TarjetaHandler) GetTarjetasByTema(title string, tema int, w http.ResponseWriter, r *http.Request) {
	tarjetas, err := h.queries.ListTarjetasByTema(context.Background(), int32(tema))
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, tarjetas, nil).Render(r.Context(), w)
}

func (h *TarjetaHandler) GetTarjetasAndTemas(title string, w http.ResponseWriter, r *http.Request) {
	tarjetas, err := h.queries.ListTarjetas(context.Background())
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	temas, err := h.queries.ListTemas(context.Background())
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, tarjetas, temas).Render(r.Context(), w)
}

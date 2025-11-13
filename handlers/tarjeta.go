package handlers

import (
	"context"
	"net/http"
	"strconv"
	db "tp5/db"
	"tp5/views"
)

type TarjetaHandler struct {
	queries *db.Queries
}

func NewTarjetaHandler(q *db.Queries) *TarjetaHandler {
	return &TarjetaHandler{queries: q}
}
func (h *TarjetaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.ListTarjetas("Welcome", w, r)
	case "/tarjetas":
		h.CreateTarjeta(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *TarjetaHandler) ListTarjetas(title string, w http.ResponseWriter, r *http.Request) {
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

	idTemaStr := r.Form.Get("tarjeta-id-tema")
	idTema, err := strconv.ParseInt(idTemaStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de tema inválido", http.StatusBadRequest)
		return
	}

	p.Pregunta = r.Form.Get("tarjeta-pregunta")
	p.Respuesta = r.Form.Get("tarjeta-respuesta")
	p.OpcionA = r.Form.Get("tarjeta-opcion-a")
	p.OpcionB = r.Form.Get("tarjeta-opcion-b")
	p.OpcionC = r.Form.Get("tarjeta-opcion-c")
	p.IDTema = int32(idTema)
	ctx := context.Background()
	h.queries.CreateTarjeta(ctx, p)
}

package handlers

import (
	"context"
	"net/http"
	"tp5/db"
	"tp5/views"
)

type TemaHandler struct {
	queries *db.Queries
}

func NewTemaHandler(q *db.Queries) *TemaHandler {
	return &TemaHandler{queries: q}
}
func (h *TemaHandler) GetTemas(title string, w http.ResponseWriter, r *http.Request) {
	tarjetas, err := h.queries.ListTemas(context.Background())
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, tarjetas, nil).Render(r.Context(), w)
}

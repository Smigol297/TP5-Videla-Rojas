package handlers

import (
	"context"
	"log"
	"net/http"
	db "tp5/db"
	"tp5/logic"
	"tp5/views"
)

func (h *TemaHandler) GetTemaByID(title string, w http.ResponseWriter, r *http.Request, id int) {
	ctx := context.Background()
	temas, err := h.queries.GetTemaById(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, temas, nil).Render(r.Context(), w)
}

func (h *TemaHandler) PutTemaByID(w http.ResponseWriter, r *http.Request, id int) {
	var p db.UpdateTemaParams
	r.ParseForm()
	p.IDTema = int32(id)
	p.NombreTema = r.Form.Get("tema-nombre")

	log.Println(p)
	// validar tema
	if err := logic.ValidateUpdateTema(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	err := h.queries.UpdateTema(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Obtener la lista actualizada (Refresh)
	temas, err := h.queries.ListTemas(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Renderizar SOLO la lista actualizada
	views.TemaList(temas).Render(r.Context(), w)
}

func (h *TemaHandler) DeleteTemaByID(w http.ResponseWriter, r *http.Request, id int) {
	ctx := context.Background()
	err := h.queries.DeleteTema(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Obtener la lista actualizada (Refresh)
	temas, err := h.queries.ListTemas(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tarjetas, err := h.queries.ListTarjetas(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	views.DeleteTemaOOB(temas, tarjetas).Render(r.Context(), w)
}

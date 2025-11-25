package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	db "tp5/db"
	"tp5/logic"
	"tp5/views"
)

func (h *TarjetaHandler) GetTarjetaByID(title string, w http.ResponseWriter, r *http.Request, id int) {
	ctx := context.Background()
	tarjetas, err := h.queries.GetTarjetaById(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, tarjetas, nil).Render(r.Context(), w)
}
func (h *TarjetaHandler) PutTarjetaByID(w http.ResponseWriter, r *http.Request, id int) {
	var p db.UpdateTarjetaParams
	p.IDTarjeta = int32(id)

	r.ParseForm()

	log.Println(p)
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
	log.Println(p)
	if err := logic.ValidateUpdateTarjeta(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	err = h.queries.UpdateTarjeta(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// En lugar de Redirect, consultamos la lista actualizada
	tarjetas, err := h.queries.ListTarjetas(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Renderizamos el componente UserList completo
	views.TarjetaTable(tarjetas).Render(r.Context(), w)
}
func (h *TarjetaHandler) DeleteTarjetaByID(w http.ResponseWriter, r *http.Request, id int) {
	ctx := context.Background()
	err := h.queries.DeleteTarjeta(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Volvemos a pedir la lista completa a la BD
	tarjetas, err := h.queries.ListTarjetas(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizamos el componente UserList completo
	views.TarjetaTable(tarjetas).Render(r.Context(), w)
}

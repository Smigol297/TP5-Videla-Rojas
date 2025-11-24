package handlers

import (
	"context"
	"net/http"
	db "tp5/db"
	"tp5/logic"
	"tp5/views"
)

func (h *UsuarioHandler) PutUsuarioByID(w http.ResponseWriter, r *http.Request, id int) {
	var p db.UpdateUsuarioParams
	p.IDUsuario = int32(id)
	r.ParseForm()
	p.NombreUsuario = r.Form.Get("usuario-nombre")
	p.Email = r.Form.Get("usuario-email")
	p.Contrasena = r.Form.Get("usuario-contrasena")
	// validar usuario
	if err := logic.ValidateUpdateUser(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	err := h.queries.UpdateUsuario(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
	// Volvemos a pedir la lista completa a la BD
	usuarios, err := h.queries.ListUsuarios(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Renderizamos el componente UserList completo
	// Esto hará que Go ejecute el "if len(usuarios) == 0" y muestre el mensaje correcto
	views.UserList(usuarios).Render(r.Context(), w)
}

func (h *UsuarioHandler) DeleteUsuarioByID(w http.ResponseWriter, r *http.Request, id int) {
	ctx := context.Background()
	err := h.queries.DeleteUsuario(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
	// Volvemos a pedir la lista completa a la BD
	usuarios, err := h.queries.ListUsuarios(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Renderizamos el componente UserList completo
	// Esto hará que Go ejecute el "if len(usuarios) == 0" y muestre el mensaje correcto
	views.UserList(usuarios).Render(r.Context(), w)
}
func (h *UsuarioHandler) GetUsuarioByID(title string, w http.ResponseWriter, r *http.Request, id int) {
	ctx := context.Background()
	usuarios, err := h.queries.GetUsuarioById(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, usuarios, nil).Render(r.Context(), w)
}

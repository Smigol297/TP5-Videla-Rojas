package handlers

import (
	"context"
	"net/http"
	"strconv"
	db "tp5ne/db/sqlc"
	"tp5ne/views"
)

type ProductHandler struct {
	queries *db.Queries
}

func NewProductHandler(q *db.Queries) *ProductHandler {
	return &ProductHandler{queries: q}
}
func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.ListProducts("Mi tienda", w, r)
	case "/products":
		h.CreateProduct(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *ProductHandler) ListProducts(title string, w http.ResponseWriter, r *http.Request) {
	products, err := h.queries.ListProducts(context.Background())
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	views.IndexPage(title, products).Render(r.Context(), w)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p db.CreateProductParams
	r.ParseForm()

	priceStr := r.Form.Get("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Precio inválido", http.StatusBadRequest)
		return
	}

	p.Name = r.Form.Get("name")
	p.Price = price
	ctx := context.Background()
	h.queries.CreateProduct(ctx, p)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

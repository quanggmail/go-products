package product

import (
	"encoding/json"
	"net/http"
)

type productCreator interface {
	Create(p *Product) error
}

type Handler struct {
	service productCreator
}

func NewHandler(service productCreator) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateProduct(
	w http.ResponseWriter,
	r *http.Request,
) {

	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.service.Create(&product)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(product)
}
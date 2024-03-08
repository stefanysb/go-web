package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"tienda/structs"
)

type MyResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MyHandler struct {
	data  []structs.Product
	lasId int
}

func NewMyHandler(products []structs.Product) *MyHandler {
	return &MyHandler{
		data: products,
	}
}

func (h *MyHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Serializar el slice de Product a JSON
		productsJSON, err := json.Marshal(h.data)
		if err != nil {
			// Manejar el error si la serialización falla
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Establecer el Content-Type a application/json
		w.Header().Set("Content-Type", "application/json")
		// Escribir el JSON en la respuesta
		w.Write(productsJSON)
	}
}

func (h *MyHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the url
		idStr := chi.URLParam(r, "productId")

		// Convertir la cadena a un entero
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// - get the user from the map
		product, ok := findProductByID(h.data, id)
		if !ok {
			code := http.StatusNotFound
			body := MyResponse{Message: "Product not found", Data: nil}
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}

		// response
		code := http.StatusOK
		body := MyResponse{Message: "Product found", Data: product}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}
}

func findProductByID(products []structs.Product, id int) (*structs.Product, bool) {
	for _, product := range products {
		if product.ID == id {
			return &product, true
		}
	}
	return nil, false
}

package main

import (
	"diseno-api/internal/handlers"
	"diseno-api/internal/repository"
	"diseno-api/internal/service"
	"diseno-api/tools"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// dependencies
	products, _ := tools.GetDataFromJSON("products.json")
	// - repository
	rp := repository.NewProductSlice(products, len(products))
	// - service
	sv := service.NewProductDefault(rp)
	// - handler
	hd := handlers.NewDefaultProduct(sv)
	// - router
	router := chi.NewRouter()
	router.Route("/products", func(r chi.Router) {
		// POST /tasks
		r.Post("/", hd.Create())
		r.Get("/", hd.Get())
		r.Get("/{productId}", hd.GetById())
		r.Get("/priceGt", hd.Search())
	})

	// server
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err)
		return
	}
}

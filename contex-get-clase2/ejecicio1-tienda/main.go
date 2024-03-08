package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"tienda/handlers"
)

func main() {

	r := chi.NewRouter()
	//obtengo los datos en forma de slice de porductos
	products, _ := GetDataFromJSON("products.json")

	//creo el handler y le paso el slice de pordcutos co el que va a trabaja
	h := handlers.NewMyHandler(products)
	r.Get("/", h.Get())
	r.Get("/product/{productId}", h.GetById())

	http.ListenAndServe(":8082", r)

}

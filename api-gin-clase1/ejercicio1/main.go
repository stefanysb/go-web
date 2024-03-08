package main

import (
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong!"))

}

func main() {
	// Crea el enrutador y registra el manejador.
	http.HandleFunc("/ping", pingHandler)

	// Inicia el servidor en el puerto 8080.
	log.Println("Server listening on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

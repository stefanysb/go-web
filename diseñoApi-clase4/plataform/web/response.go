package web

import (
	"encoding/json"
	"net/http"
)

// Text escribe una respuesta de texto plano al cliente.
func Text(w http.ResponseWriter, code int, body string) {
	// Establece el tipo de contenido de la respuesta a texto plano.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Establece el código de estado HTTP de la respuesta.
	w.WriteHeader(code)

	// Escribe el cuerpo de texto plano en la respuesta.
	w.Write([]byte(body))
}

// JSON escribe una respuesta en formato JSON al cliente.
func JSON(w http.ResponseWriter, code int, body any) {
	// Si el cuerpo es nil, solo establece el código de estado y termina.
	if body == nil {
		w.WriteHeader(code)
		return
	}

	// Intenta serializar el cuerpo a JSON.
	bytes, err := json.Marshal(body)
	if err != nil {
		// Si hay error en la serialización, envía un error 500.
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Establece el tipo de contenido de la respuesta a JSON.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Establece el código de estado HTTP de la respuesta.
	w.WriteHeader(code)

	// Escribe el cuerpo JSON serializado en la respuesta.
	w.Write(bytes)
}

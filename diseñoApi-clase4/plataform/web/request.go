package web

import (
	"encoding/json" // Importa el paquete para decodificar JSON.
	"errors"        // Importa el paquete para manejar errores.
	"fmt"           // Importa el paquete para operaciones de entrada/salida formateada.
	"net/http"      // Importa el paquete para manejar solicitudes HTTP.
	"regexp"        // Importa el paquete para trabajar con expresiones regulares.
	"strings"       // Importa el paquete para operaciones con strings.
)

// Define un error para cuando el JSON de la solicitud es inválido.
var (
	ErrRequestJSONInvalid = errors.New("request json invalid")
)

// JSON decodifica el JSON del cuerpo de la solicitud HTTP y lo almacena en la variable apuntada por ptr.
func JSONdecoder(r *http.Request, ptr any) (err error) {
	// Intenta decodificar el cuerpo de la solicitud.
	err = json.NewDecoder(r.Body).Decode(ptr)
	if err != nil {
		// Si hay un error, envuelve el error original con ErrRequestJSONInvalid para más contexto.
		err = fmt.Errorf("%w. %v", ErrRequestJSONInvalid, err)
		return
	}

	return
}

// Define un error para cuando el parámetro de la ruta de la solicitud es inválido.
var (
	ErrRequestPathParamInvalid = errors.New("request path param invalid")
)

// PathLastParam retorna el valor del último parámetro en la ruta de la solicitud HTTP.
func PathLastParam(r *http.Request) (value string, err error) {
	// Obtiene la ruta de la URL de la solicitud.
	path := r.URL.Path

	// Define una expresión regular para validar y extraer el último parámetro de la ruta.
	rx := regexp.MustCompile(`^/(.*/)*([0-9a-zA-Z]+)$`)
	if !rx.MatchString(path) {
		// Si la ruta no coincide con el patrón de la expresión regular, retorna un error.
		err = ErrRequestPathParamInvalid
		return
	}

	// Divide la ruta en segmentos usando el carácter `/`.
	sl := strings.Split(path, "/")
	// Extrae el último segmento de la ruta, que se asume es el parámetro deseado.
	value = sl[len(sl)-1]
	return
}

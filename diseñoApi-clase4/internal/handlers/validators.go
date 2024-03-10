package handlers

import (
	"regexp"
	"time"
)

// validarFecha verifica si una cadena de texto es una fecha válida en el formato DD/MM/YYYY.
// Retorna true si la fecha es válida, false en caso contrario.
func validateDate(fecha string) bool {
	// Primero, verifica el formato usando una expresión regular.
	// La expresión regular busca un patrón de cuatro dígitos, seguido por un slash,
	// luego dos dígitos, otro slash y finalmente cuatro dígitos.
	match, _ := regexp.MatchString(`^\d{2}/\d{2}/\d{4}$`, fecha)
	if !match {
		return false
	}

	// Intenta parsear la cadena a una fecha usando el formato DD/MM/YYYY.
	// time.Parse retorna un error si la cadena no se puede parsear a una fecha válida.
	_, err := time.Parse("02/01/2006", fecha)
	if err != nil {
		return false // La cadena no es una fecha válida.
	}

	// Si no hay errores, la fecha es válida.
	return true
}

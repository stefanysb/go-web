package tools

import (
	"diseno-api"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetDataFromJSON(path string) ([]internal.Product, error) {
	// Abrir el archivo JSON
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	// Leer el contenido del archivo JSON
	jsonData, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading JSON data: %v\n", err)
		return nil, err
	}

	// Deserializar el JSON directamente en un slice de Product
	var products []internal.Product
	err = json.Unmarshal(jsonData, &products)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return nil, err
	}

	return products, nil
}

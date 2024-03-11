package internal

import (
	"errors"
)

// Product define la estructura para un producto
type Product struct {
	ID          int     `json:"id,omitempty"` // ID es el identificador único del producto
	Name        string  `json:"name"`         // Name es el nombre del producto
	Quantity    int     `json:"quantity"`     // Quantity indica la cantidad disponible del producto
	CodeValue   string  `json:"code_value"`   // CodeValue es el código de valor único para el producto
	IsPublished bool    `json:"is_published"` // IsPublished indica si el producto está publicado
	Expiration  string  `json:"expiration"`   // Expiration es la fecha de vencimiento del producto
	Price       float64 `json:"price"`        // Price es el precio del producto
}

var (
	// ErrProductNotFound se utiliza cuando un producto no se encuentra
	ErrProductNotFound = errors.New("product not found")
	// ErrProductDuplicated se utiliza cuando un producto ya existe
	ErrProductDuplicated = errors.New("product already exists")
	// ErrProductInvalidField se utiliza cuando un producto tiene campos inválidos
	ErrProductInvalidField = errors.New("product has invalid fields")
	///ErrProductInternal se utiliza cuando un producto no puede ser procesado debido a un error interno
	ErrProductInternal = errors.New("product can't be processed")
)

// ProductRepository es una interfaz para operaciones de base de datos relacionadas con productos
type ProductRepository interface {
	// Save persiste un producto
	Save(product *Product) error
	// FindByID encuentra un producto por su ID
	FindByID(id int) (*Product, error)
	// GetAll encuentra todos los porductos de inventario
	GetAll() []Product
	// encuentra todos los porductos cuyo precio sea mayor al apsado por parámetro
	GetByQuery(price float64) []Product
	// actualiza un prodcuto
	Update(product Product) error
	//reliza una actilizacion parcial de los campos
	UpdatePartial(id int, fields map[string]any) (err error)
}

// ProductService es una interfaz para operaciones de servicio relacionadas con productos
type ProductService interface {
	// Save persiste un producto a través de la capa de servicio
	Save(product *Product) error
	FindByID(id int) (*Product, error)
	// GetAll encuentra todos los porductos de inventario
	GetAll() []Product
	// encuentra todos los porductos cuyo precio sea mayor al apsado por parámetro
	GetByQuery(price float64) []Product
	// actualiza un prodcuto
	Update(product Product) error
	//reliza una actilizacion parcial de los campos
	UpdatePartial(id int, fields map[string]any) (err error)
}

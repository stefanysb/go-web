package repository

import (
	"diseno-api"
	"diseno-api/tools"
	"fmt"
)

// NewProductSlice inicializa una nueva instancia de ProductSlice.
// Toma un producto (o un slice de productos, dependiendo de cómo quieras inicializarlo) y un lastID como argumentos.
func NewProductSlice(db []internal.Product, lastID int) *ProductSlice {
	// Valores predeterminados
	if db == nil {
		// Si db es nil, inicializa un slice vacío de Product.
		db = make([]internal.Product, 0)
	}

	// Crea la instancia de ProductSlice con el slice de productos y el último ID proporcionado.
	return &ProductSlice{
		db:     db,
		lastID: lastID,
	}
}

// ProductSlice es una implementación que permite almacenar y gestionar una colección de productos.
type ProductSlice struct {
	db     []internal.Product
	lastID int
}

// Save agrega un nuevo producto al slice ProductSlice si no existe un producto con el mismo nombre.
func (p *ProductSlice) Save(product *internal.Product) error {
	// Verifica si el producto ya existe basado en un atributo único, como podría ser el ID.
	for _, prod := range p.db {
		if prod.CodeValue == product.CodeValue {
			// Retorna un error específico si el producto ya existe.
			return internal.ErrProductDuplicated
		}
	}

	// Incrementa lastID para asignar un nuevo identificador único al producto.
	p.lastID++

	// Asigna el nuevo ID al producto.
	product.ID = p.lastID

	// Añade el producto al slice.
	p.db = append(p.db, *product)

	// Si todo es exitoso, retorna nil indicando que no hubo errores.
	return nil
}

// FindByID encuentra un producto por su ID.
func (p *ProductSlice) FindByID(id int) (*internal.Product, error) {
	for i := range p.db {
		if p.db[i].ID == id {
			return &p.db[i], nil // Devuelve la referencia al elemento dentro del slice
		}
	}
	return nil, internal.ErrProductNotFound
}
func (p *ProductSlice) GetAll() []internal.Product {
	// Inicializa un nuevo slice para los productos copiados con la misma longitud que p.db.
	copiedProducts := make([]internal.Product, len(p.db))

	// Usa copy para copiar los elementos de p.db a copiedProducts.
	copy(copiedProducts, p.db)

	// Retorna el nuevo slice, que es una copia de p.db.
	return copiedProducts
}

func (p *ProductSlice) GetByQuery(price float64) []internal.Product {
	var result []internal.Product // Creamos una lista vacía para almacenar los productos que cumplan con la condición.

	for _, product := range p.db { // Recorremos cada producto en la lista de productos.
		if product.Price > price { // Comprobamos si el precio del producto es mayor al precio especificado.
			result = append(result, product) // Si cumple con la condición, añadimos el producto a la lista de resultados.
		}
	}

	return result // Retornamos la lista de productos que cumplen con la condición y nil como error.
}

// Update actualiza un producto en ProductSlice.
func (p *ProductSlice) Update(product internal.Product) error {
	// valido que el codigo no esté duplicado
	for _, pro := range p.db {
		if pro.CodeValue == product.CodeValue {
			// Actualiza directamente el producto en el slice.

			return internal.ErrProductDuplicated
		}
	}
	for i, pro := range p.db {
		if pro.ID == product.ID {
			// Actualiza directamente el producto en el slice.
			p.db[i] = product
			return nil
		}
	}
	return internal.ErrProductNotFound
}

func (p *ProductSlice) UpdatePartial(id int, fields map[string]any) (err error) {
	// busco el producto
	prod, err := p.FindByID(id)
	if err != nil {
		return internal.ErrProductNotFound
	}

	//iterp spbre los campos para actualizar, Si el campo existe en el mapa de campos, lo actulizo
	// es decir, si el campo viene entre los capos de la peticion, lo actulizo
	for key, value := range fields {
		switch key {

		case "Name", "name":
			name, ok := value.(string)
			if !ok {
				return &tools.FieldError{
					Field: key,
					Msg:   fmt.Sprintf("%s error type assertion", key),
				}
			}
			prod.Name = name
		case "Quantity", "quantity":
			floatVal, ok := value.(float64)
			if !ok {
				return &tools.FieldError{
					Field: key,
					Msg:   fmt.Sprintf("%s error type assertion\"", key),
				}

			}
			prod.Quantity = int(floatVal)
		case "CodeValue", "code_value":
			codeValue, ok := value.(string)
			if !ok {
				return &tools.FieldError{
					Field: key,
					Msg:   fmt.Sprintf("%s error type assertion\"", key),
				}
			}
			prod.CodeValue = codeValue
		case "IsPublished", "is_published":
			isPublished, ok := value.(bool)
			if !ok {
				return &tools.FieldError{
					Field: key,
					Msg:   fmt.Sprintf("%s error type assertion\"", key),
				}
			}
			prod.IsPublished = isPublished
		case "Expiration", "expiration":
			expiration, ok := value.(string)
			if !ok {
				return &tools.FieldError{
					Field: key,
					Msg:   fmt.Sprintf("%s error type assertion\"", key),
				}
			}
			prod.Expiration = expiration
		case "Price", "price":
			price, ok := value.(float64)
			if !ok {
				return &tools.FieldError{
					Field: key,
					Msg:   fmt.Sprintf("%s error type assertion\"", key),
				}
			}
			prod.Price = price
		default:
			return &tools.FieldError{
				Field: "default",
				Msg:   fmt.Sprintf("%s error type assertion\"", key),
			}
		}
	}

	return nil
	// actulizo la teraa
}

// Paquete service gestiona la lógica de servicios para productos.
package service

// Importa funcionalidades del paquete interno diseno-api.
import (
	internal "diseno-api"
)

// NewProductDefault crea una instancia de ProductDefault con el repositorio dado.
func NewProductDefault(repository internal.ProductRepository) *ProductDefault {
	// Inicializa ProductDefault con un repositorio específico.
	return &ProductDefault{
		repository: repository,
	}
}

// ProductDefault maneja operaciones de servicio para productos.
type ProductDefault struct {
	repository internal.ProductRepository // Referencia al repositorio para operaciones de producto.
}

// Save guarda un producto usando el repositorio asociado.
func (p *ProductDefault) Save(product *internal.Product) (err error) {
	// Delega la operación de guardado al repositorio y retorna el resultado.
	err = p.repository.Save(product)
	return
}

func (p *ProductDefault) GetAll() []internal.Product {
	// Delega la operación de guardado al repositorio y retorna el resultado.

	return p.repository.GetAll()
}

func (p *ProductDefault) FindByID(id int) (*internal.Product, error) {
	//tomamos el id

	// lo buscamos en el repository, si no existe retornamos el error
	product, err := p.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// devolvemos el porducto
	return product, nil

}

// encuentra todos los porductos cuyo precio sea mayor al apsado por parámetro
func (p *ProductDefault) GetByQuery(price float64) []internal.Product {

	result := p.repository.GetByQuery(price)
	return result
}

func (p *ProductDefault) Update(product internal.Product) error {
	result := p.repository.Update(product)
	return result
}

func (p *ProductDefault) UpdatePartial(id int, fields map[string]interface{}) (err error) {
	// update the task partially
	err = p.repository.UpdatePartial(id, fields)
	return
}

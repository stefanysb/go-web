package handlers

import (
	internal "diseno-api"
	"diseno-api/plataform/web"
	"diseno-api/tools"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

func NewDefaultProduct(service internal.ProductService) *DefaultProduct {
	return &DefaultProduct{
		service: service,
	}

}

// la estrutura por default que usara el handler
type DefaultProduct struct {
	service internal.ProductService
}

type ProductJson struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// la manera en que un porducto viene de una request
type ProductJsonBody struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// CreateProduct es un método que crea un nuevo producto. Se implementa como un manejador HTTP.
func (d *DefaultProduct) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extraigo el body como una cadena de bytess
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			web.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request body"})
			return
		}

		// lo convierto en un mapa para mirar si me faltan daors
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			web.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		// valido si me faltan datos
		if err := tools.CheckFieldExistance(bodyMap, "name", "quantity", "code_value", "is_published", "expiration", "price"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				web.JSON(w, http.StatusBadRequest, map[string]any{"message": fmt.Sprintf("%s is required", fieldError.Field)})
				return
			}

			web.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			return
		}

		// lo convierto es una estrutura de respuesta, si no lo puedo convertir digo el envio la respuesta que i
		var body ProductJsonBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			web.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		// serilizo de JsonBpdy a Porducr

		// Serializar la solicitud en un producto (paso 1)
		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		// Guardar el producto (paso 2)
		err1 := d.service.Save(&product)
		if err1 != nil {

			// Manejo de errores al guardar el producto
			// Puedes personalizar la respuesta basándote en el tipo de error
			// pendiente pregunta porque en lo case hat varios erroes, pero miradno solo me aprece uno
			web.JSON(w, http.StatusInternalServerError, map[string]any{"message": "failed to save product"})
			return
		}

		// Responder al cliente (paso 3)
		// Puedes ajustar los campos de la respuesta según tu estructura de Product
		responseProduct := ProductJson{
			Id:          product.ID, // Asegúrate de que el Id se establece correctamente después de guardar
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		web.JSON(w, http.StatusCreated, map[string]any{"message": "product created successfully", "product": responseProduct})
	}

}

func (d *DefaultProduct) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Serializar el slice de Product a JSON
		productos := d.service.GetAll()
		_, err := json.Marshal(&productos)
		if err != nil {
			// Manejar el error si la serialización falla
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// envio respuesta
		web.JSON(w, http.StatusOK, &productos)
		return
	}
}

func (d *DefaultProduct) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//sacar el id de la url
		idStr := chi.URLParam(r, "productId")

		//lo converito a entero porque el método lo pide en entero
		idProduct, err := strconv.Atoi(idStr)
		if err != nil {
			web.JSON(w, http.StatusBadRequest, "invalid id")
		}

		//busco y retorno el porcutp
		productResponse, err := d.service.FindByID(idProduct)

		if errors.Is(err, internal.ErrProductNotFound) {
			web.Text(w, http.StatusNotFound, "product not found ")
			return
		}
		web.JSON(w, http.StatusOK, productResponse)

		return
	}
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/henrymoreirasilva/go-api/internal/dto"
	"github.com/henrymoreirasilva/go-api/internal/entity"
	"github.com/henrymoreirasilva/go-api/internal/infra/database"
	entityPkg "github.com/henrymoreirasilva/go-api/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewHandlerProduct(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create product
// @Summary			Create product
// @Description		Create product
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param 			request	body	dto.CreateProductInput	true	"Product request"
// @Success 		201
// @Failure			400	{object}	Error
// @Failure 		500	{object}	Error
// @Router			/products	[post]
// @Security		ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product dto.CreateProductInput

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if err = h.ProductDB.Create(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get product
// @Summary			Get product
// @Description		Get product
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param 			id	path	string	true	"Product ID" Format(uuid)
// @Success 		200	{object}	entity.Product
// @Failure			400	{object}	Error
// @Failure 		500	{object}	Error
// @Router			/products/{id}	[get]
// @Security		ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

// Update product
// @Summary			Update product
// @Description		Update product
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param 			id	path	string true	"Product ID" Format(uuid)
// @Param 			request	body	dto.CreateProductInput true "Product request"
// @Success 		200
// @Failure			400	{object}	Error
// @Failure			500	{object}	Error
// @Router			/products/{id}	[put]
// @Security		ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// fa64dffb-0074-46b9-b200-f166c780cb38
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if err = h.ProductDB.Update(&product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		http.Error(w, "delete error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get products
// @Summary			Get products
// @Description		Get products
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param 			page	query	string	false	"Page number"
// @Param 			limit	query	string	false	"Limit number"
// @Success 		200	{array}	entity.Product
// @Failure			400
// @Router			/products	[get]
// @Security		ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-type:", "pplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}

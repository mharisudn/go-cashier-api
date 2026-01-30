package handlers

import (
	"cashier-api/helpers"
	"cashier-api/models"
	"cashier-api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// HandleProducts - GET /api/products
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		helpers.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		helpers.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, "Products retrieved successfully", products)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.ProductCreate
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		helpers.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		helpers.JSONError(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, "Product retrieved successfully", product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product models.ProductUpdate
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, "Product updated successfully", product)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		helpers.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, "Product deleted successfully", nil)
}

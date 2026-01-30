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

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategories - GET /api/categories
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		helpers.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		helpers.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, "Categories retrieved successfully", categories)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, "Category created successfully", category)
}

func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.service.GetByIDWithProducts(id)
	if err != nil {
		helpers.JSONError(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, "Category retrieved successfully", category)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, "Category updated successfully", category)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		helpers.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, "Category deleted successfully", nil)
}

package main

import (
	"cashier-api/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var categories = []model.Category{
	{
		ID:          1,
		Name:        "Food",
		Description: "Food category",
	},
	{
		ID:          2,
		Name:        "Drink",
		Description: "Drink category",
	},
	{
		ID:          3,
		Name:        "Snack",
		Description: "Snack category",
	},
}

var products = []model.Product{
	{
		ID:          1,
		Name:        "Chicken",
		Description: "Chicken original",
		Price:       10000,
		CategoryID:  1,
	},
	{
		ID:          2,
		Name:        "Beef",
		Description: "Beef original",
		Price:       20000,
		CategoryID:  1,
	},
	{
		ID:          3,
		Name:        "Fish",
		Description: "Fish original",
		Price:       30000,
		CategoryID:  1,
	},
}

// get category by ID
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, cat := range categories {
		if cat.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cat)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// get product by ID
func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, prod := range products {
		if prod.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(prod)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// update category by ID
func updateCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var category model.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, cat := range categories {
		if cat.ID == id {
			category.ID = id
			categories[i] = category
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// update product by ID
func updateProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product model.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, prod := range products {
		if prod.ID == id {
			product.ID = id
			products[i] = product
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// delete category by ID
func deleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, cat := range categories {
		if cat.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted"})
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// delete product by ID
func deleteProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, prod := range products {
		if prod.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	/*
		Handle category requests
		- GET /api/categories/{id}
		- PUT /api/categories/{id}
		- DELETE /api/categories/{id}
	*/
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCategoryByID(w, r)
		case http.MethodPut:
			updateCategoryByID(w, r)
		case http.MethodDelete:
			deleteCategoryByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	/*
		Handle category requests
		- GET /api/categories
		- POST /api/categories
	*/
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
			return
		case http.MethodPost:
			var category model.Category
			err := json.NewDecoder(r.Body).Decode(&category)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			category.ID = len(categories) + 1
			categories = append(categories, category)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(category)
			return
		}
	})

	/*
		Handle product requests
		- GET /api/products/{id}
		- PUT /api/products/{id}
		- DELETE /api/products/{id}
	*/
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProductByID(w, r)
		case http.MethodPut:
			updateProductByID(w, r)
		case http.MethodDelete:
			deleteProductByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	/*
		Handle product requests
		- GET /api/products
		- POST /api/products
	*/
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
			return
		case http.MethodPost:
			var product model.Product
			err := json.NewDecoder(r.Body).Decode(&product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			product.ID = len(products) + 1
			products = append(products, product)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
			return
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	port := ":8080"

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}

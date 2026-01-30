package repositories

import (
	"cashier-api/models"
	"database/sql"
	"errors"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// Get All Products
func (repo *CategoryRepository) GetAll() ([]models.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.Name, &cat.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

// Create Category
func (repo *CategoryRepository) Create(category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	return err
}

// Get Category By ID
func (repo *CategoryRepository) GetByID(id int) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	var cat models.Category
	err := repo.db.QueryRow(query, id).Scan(&cat.ID, &cat.Name, &cat.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("Category not found")
	}

	if err != nil {
		return nil, err
	}
	return &cat, nil
}

// Get Category By ID with Products
func (repo *CategoryRepository) GetByIDWithProducts(id int) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"
	var cat models.Category
	err := repo.db.QueryRow(query, id).Scan(&cat.ID, &cat.Name, &cat.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("Category not found")
	}

	if err != nil {
		return nil, err
	}

	// Get Products by Category ID
	queryProd := "SELECT id, name, price, stock FROM products WHERE category_id = $1"
	rows, err := repo.db.Query(queryProd, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cat.Products = make([]models.ProductList, 0)
	for rows.Next() {
		var prod models.ProductList
		err := rows.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Stock)
		if err != nil {
			return nil, err
		}
		cat.Products = append(cat.Products, prod)
	}
	return &cat, nil
}

// Update Category
func (repo *CategoryRepository) Update(category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	result, err := repo.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Category not found")
	}
	return nil
}

// Delete Category
func (repo *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Category not found")
	}
	return nil
}

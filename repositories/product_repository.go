package repositories

import (
	"cashier-api/models"
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	query := `
	SELECT p.id, p.name, p.price, p.stock, p.category_id,
	COALESCE(c.name, '') as cat_name,
	COALESCE(c.description, '') as cat_description
	FROM products p
	LEFT JOIN categories c ON p.category_id = c.id
	`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var prod models.Product

		// Initialize Struct Category Detail
		prod.Category = &models.CategoryDetail{}

		err := rows.Scan(
			&prod.ID,
			&prod.Name,
			&prod.Price,
			&prod.Stock,
			&prod.CategoryID,
			&prod.Category.Name,
			&prod.Category.Description,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	// Query
	query := `
	SELECT p.id, p.name, p.price, p.stock, p.category_id,
	COALESCE(c.name, '') as cat_name,
	COALESCE(c.description, '') as cat_description
	FROM products p
	LEFT JOIN categories c ON p.category_id = c.id
	WHERE p.id = $1
	`
	// Temporary Buffer Product
	var prod models.Product
	// Initialize Struct Category Detail
	prod.Category = &models.CategoryDetail{}

	err := repo.db.QueryRow(query, id).Scan(
		&prod.ID,
		&prod.Name,
		&prod.Price,
		&prod.Stock,
		&prod.CategoryID,
		&prod.Category.Name,
		&prod.Category.Description,
	)
	// Check Available Product
	if err == sql.ErrNoRows {
		return nil, errors.New("Product not found")
	}
	// Check Error
	if err != nil {
		return nil, err
	}
	// Return Product
	return &prod, nil
}

func (repo *ProductRepository) Update(product *models.ProductUpdate) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	if err != nil {
		return err
	}
	// Check Affected Rows
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Check Affected Rows
	if rows == 0 {
		return errors.New("Product not found")
	}
	// Return nil if success
	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Check Affected Rows
	if rows == 0 {
		return errors.New("Product not found")
	}
	// Return nil if success
	return nil
}

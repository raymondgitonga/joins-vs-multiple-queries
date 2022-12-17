package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/raymondgitonga/joins-vs-multiple-queries/internal/core/dormain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) (*ProductRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is null")
	}
	return &ProductRepository{
		db: db,
	}, nil
}

func (r *ProductRepository) GetProductJoin(ctx context.Context, productID int) (*dormain.Product, error) {
	product := &dormain.Product{}
	query := `select product_catalog.name, product_catalog.price, product_quantity.quantity, product_category.name
              from product_catalog join product_quantity on product_catalog.id = product_quantity.catalog_id
              join product_category on product_catalog.category_id = product_category.id where product_catalog.id = $1;`

	row := r.db.QueryRowContext(ctx, query, productID)

	err := row.Scan(&product.Name, &product.Price, &product.Quantity, &product.Category)
	if err != nil {
		return nil, fmt.Errorf("error fetching product, %w", err)
	}

	return product, nil
}

func (r *ProductRepository) GetProduct(ctx context.Context, productID int) (*dormain.Product, error) {
	product := &dormain.Product{}
	query := `select product_catalog.name, product_catalog.price from product_catalog where product_catalog.id = $1;`

	row := r.db.QueryRowContext(ctx, query, productID)

	err := row.Scan(&product.Name, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("error fetching product, %w", err)
	}

	return product, nil
}

func (r *ProductRepository) GetProductQuantity(ctx context.Context, productID int) (int, error) {
	product := &dormain.Product{}
	query := `select product_quantity.quantity from product_quantity where product_quantity.catalog_id = $1;`

	row := r.db.QueryRowContext(ctx, query, productID)

	err := row.Scan(&product.Quantity)
	if err != nil {
		return 0, fmt.Errorf("error fetching product quantity, %w", err)
	}

	return product.Quantity, nil
}

func (r *ProductRepository) GetProductCategory(ctx context.Context, productID int) (int, error) {
	product := &dormain.Product{}
	query := `select product_catalog_category.category_id from product_catalog_category where product_catalog_category.catalog_id = $1;`

	row := r.db.QueryRowContext(ctx, query, productID)

	err := row.Scan(&product.CategoryID)
	if err != nil {
		return 0, fmt.Errorf("error fetching product quantity, %w", err)
	}

	return product.CategoryID, nil
}

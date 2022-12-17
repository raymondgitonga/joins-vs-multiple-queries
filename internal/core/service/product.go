package service

import (
	"context"
	"github.com/raymondgitonga/joins-vs-multiple-queries/internal/core/dormain"
)

type ProductRepository interface {
	GetProductJoin(ctx context.Context, productID int) (*dormain.Product, error)
	GetProduct(ctx context.Context, productID int) (*dormain.Product, error)
	GetProductQuantity(ctx context.Context, productID int) (int, error)
	GetProductCategory(ctx context.Context, productID int) (int, error)
}

type ProductService struct {
	productRepo ProductRepository
}

func NewProductService(productRepo ProductRepository) *ProductService {
	return &ProductService{productRepo}
}

func (s *ProductService) GetProductJoin(ctx context.Context, productID int) (*dormain.Product, error) {
	product, err := s.productRepo.GetProductJoin(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, productID int) (*dormain.Product, error) {
	product, err := s.productRepo.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	quantity, err := s.productRepo.GetProductQuantity(ctx, productID)
	if err != nil {
		return nil, err
	}

	category, err := s.productRepo.GetProductQuantity(ctx, productID)
	if err != nil {
		return nil, err
	}

	product.Quantity = quantity
	product.CategoryID = category

	return product, nil
}

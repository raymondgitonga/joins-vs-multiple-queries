package service

import (
	"context"
	"github.com/raymondgitonga/joins-vs-multiple-queries/internal/core/dormain"
)

type FetchProduct struct {
	product *dormain.Product
	err     error
}

type FetchQuantity struct {
	quantity int
	err      error
}

type FetchCategory struct {
	category int
	err      error
}

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

func (s *ProductService) GetProductSync(ctx context.Context, productID int) (*dormain.Product, error) {
	product, err := s.productRepo.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	quantity, err := s.productRepo.GetProductQuantity(ctx, productID)
	if err != nil {
		return nil, err
	}

	category, err := s.productRepo.GetProductCategory(ctx, productID)
	if err != nil {
		return nil, err
	}

	product.Quantity = quantity
	product.CategoryID = category

	return product, nil
}

func (s *ProductService) GetProductAsync(ctx context.Context, productID int) (*dormain.Product, error) {
	productChan := make(chan FetchProduct)
	quantityChan := make(chan FetchQuantity)
	categoryChan := make(chan FetchCategory)

	go s.fetchProduct(ctx, productID, productChan)
	go s.fetchProductQuantity(ctx, productID, quantityChan)
	go s.fetchProductCategory(ctx, productID, categoryChan)

	fetchProduct := <-productChan
	if fetchProduct.err != nil {
		return nil, fetchProduct.err
	}

	fetchQuantity := <-quantityChan
	if fetchQuantity.err != nil {
		return nil, fetchQuantity.err
	}

	fetchCategory := <-categoryChan
	if fetchCategory.err != nil {
		return nil, fetchCategory.err
	}

	product := fetchProduct.product
	product.Quantity = fetchQuantity.quantity
	product.CategoryID = fetchCategory.category

	return product, nil
}

func (s *ProductService) fetchProduct(ctx context.Context, productID int, productChan chan FetchProduct) {
	product, err := s.productRepo.GetProduct(ctx, productID)
	if err != nil {
		productChan <- FetchProduct{product: nil, err: err}
		return
	}

	productChan <- FetchProduct{product: product, err: nil}
}

func (s *ProductService) fetchProductQuantity(ctx context.Context, productID int, quantityChan chan FetchQuantity) {
	quantity, err := s.productRepo.GetProductQuantity(ctx, productID)
	if err != nil {
		quantityChan <- FetchQuantity{quantity: 0, err: err}
		return
	}

	quantityChan <- FetchQuantity{quantity: quantity, err: nil}
}

func (s *ProductService) fetchProductCategory(ctx context.Context, productID int, categoryChan chan FetchCategory) {
	category, err := s.productRepo.GetProductCategory(ctx, productID)
	if err != nil {
		categoryChan <- FetchCategory{category: 0, err: err}
		return
	}

	categoryChan <- FetchCategory{category: category, err: nil}
}

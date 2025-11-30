package service

import (
	"e-commerce/model"
	"e-commerce/repository"
	"errors"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProduct(id int64) (*model.Product, error)
	ListProducts() ([]model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int64) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *model.Product) error {
	if product.Name == "" || product.Price < 0 {
		return errors.New("invalid product data")
	}
	return s.repo.Create(product)
}

func (s *productService) GetProduct(id int64) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) ListProducts() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) UpdateProduct(product *model.Product) error {
	if product.ID == 0 {
		return errors.New("product id required")
	}
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id int64) error {
	return s.repo.Delete(id)
}

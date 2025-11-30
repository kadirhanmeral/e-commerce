package service

import (
	"e-commerce/model"
	"e-commerce/repository"
	"errors"
)

type CartService interface {
	AddToCart(userID, productID int64, quantity int) error
	GetCart(userID int64) (*model.Cart, error)
	ClearCart(userID int64) error
}

type cartService struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewCartService(cartRepo repository.CartRepository, productRepo repository.ProductRepository) CartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *cartService) AddToCart(userID, productID int64, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	product, err := s.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	item := &model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     product.Price,
	}

	return s.cartRepo.AddItem(item)
}

func (s *cartService) GetCart(userID int64) (*model.Cart, error) {
	items, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return nil, err
	}

	cart := &model.Cart{
		UserID: userID,
		Items:  items,
		Total:  0,
	}

	for _, item := range items {
		cart.Total += item.Price * float64(item.Quantity)
	}

	return cart, nil
}

func (s *cartService) ClearCart(userID int64) error {
	return s.cartRepo.ClearCart(userID)
}

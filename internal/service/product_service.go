package service

import (
	"github.com/destafajri/golang-fiber/internal/model"
	"github.com/destafajri/golang-fiber/internal/repository"
)

type ProductService interface {
	Create(request model.CreateProductRequest) (response model.CreateProductResponse)
	List() (responses []model.GetProductResponse)
}

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
	}
}
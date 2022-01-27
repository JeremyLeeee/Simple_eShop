package servies

import (
	"eshop/datamodels"
	"eshop/repositories"
)

type IProductService interface {
	GetProductById(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductById(int64) bool
	InsertProduct(*datamodels.Product) (int64, error)
	UpdateProduct(*datamodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

func NewProductService(productRepository repositories.IProduct) IProductService {
	return &ProductService{productRepository}
}

func (s *ProductService) GetProductById(id int64) (*datamodels.Product, error) {
	return s.productRepository.SelectByKey(id)
}

func (s *ProductService) GetAllProduct() ([]*datamodels.Product, error) {
	return s.productRepository.SelectAll()
}

func (s *ProductService) DeleteProductById(id int64) bool {
	return s.productRepository.Delete(id)
}

func (s *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return s.productRepository.Insert(product)
}

func (s *ProductService) UpdateProduct(product *datamodels.Product) error {
	return s.productRepository.Update(product)
}

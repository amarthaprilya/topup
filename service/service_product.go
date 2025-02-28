package service

import (
	"camera-rent/entity"
	"camera-rent/input"
	"camera-rent/repository"
	"errors"
)

type ServiceProduct interface {
	CreateProduct(input input.ProductInput) (*entity.Products, error)
	GetProducts() ([]*entity.Products, error)
	GetProduct(ID int) (*entity.Products, error)
	DeleteProduct(ID int) error
	UpdateProduct(ID int, input input.ProductInput) (*entity.Products, error)
}

type serviceProduct struct {
	repositoryProduct  repository.RepositoryProduct
	repositoryCategory repository.RepositoryCategory
}

func NewServiceProduct(repositoryProduct repository.RepositoryProduct, repositoryCategory repository.RepositoryCategory) *serviceProduct {
	return &serviceProduct{repositoryProduct, repositoryCategory}
}

func (s *serviceProduct) GetProductByCategory(ID int) ([]*entity.Products, error) {
	product, err := s.repositoryProduct.FindAllProductByCategory(ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) GetProducts() ([]*entity.Products, error) {

	product, err := s.repositoryProduct.FindAll()
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) UpdateProduct(ID int, input input.ProductInput) (*entity.Products, error) {
	find, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return find, err
	}

	find.Name = input.Name
	find.RentCost = input.RentCost
	find.Stock = input.Stock

	newProduct, err := s.repositoryProduct.Update(find)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *serviceProduct) CreateProduct(input input.ProductInput) (*entity.Products, error) {

	if input.CategoryID == 0 {
		return nil, errors.New("CategoryID tidak boleh kosong")
	}

	// Cek apakah kategori ada di database
	_, err := s.repositoryCategory.FindById(input.CategoryID)
	if err != nil {
		return nil, errors.New("kategori tidak ditemukan")
	}

	product := &entity.Products{}

	product.Name = input.Name
	product.RentCost = input.RentCost
	product.Stock = input.Stock
	product.Description = input.Description
	product.CategoryID = input.CategoryID

	newProduct, err := s.repositoryProduct.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *serviceProduct) GetProduct(ID int) (*entity.Products, error) {

	product, err := s.repositoryProduct.FindById(ID)

	if err != nil {
		return nil, err
	}

	if product.ID == 0 {
		return nil, err
	}

	return product, nil
}

func (s *serviceProduct) DeleteProduct(ID int) error {
	// Temukan produk berdasarkan ID
	product, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return err
	}

	// Hapus produk dari basis data
	_, err = s.repositoryProduct.Delete(product)
	if err != nil {
		return err
	}

	return nil
}

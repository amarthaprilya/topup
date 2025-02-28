package service

import (
	"camera-rent/entity"
	"camera-rent/input"
	"camera-rent/repository"
)

type ServiceCategory interface {
	CreateCategory(input input.CategoryInput) (*entity.Category, error)
	GetCategorys() ([]*entity.Category, error)
	GetCategoryByID(ID int) (*entity.Category, error)
	DeleteCategory(ID int) error
	UpdateCategory(ID int, input input.CategoryInput) (*entity.Category, error)
}

type serviceCategory struct {
	repositoryCategory repository.RepositoryCategory
}

func NewServiceCategory(repositoryCategory repository.RepositoryCategory) *serviceCategory {
	return &serviceCategory{repositoryCategory}
}

func (s *serviceCategory) GetCategorys() ([]*entity.Category, error) {

	product, err := s.repositoryCategory.FindAll()
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceCategory) UpdateCategory(ID int, input input.CategoryInput) (*entity.Category, error) {
	find, err := s.repositoryCategory.FindById(ID)
	if err != nil {
		return find, err
	}

	find.Name = input.Name

	newProduct, err := s.repositoryCategory.Update(find)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *serviceCategory) CreateCategory(input input.CategoryInput) (*entity.Category, error) {
	product := &entity.Category{}

	product.Name = input.Name

	newProduct, err := s.repositoryCategory.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *serviceCategory) GetCategoryByID(ID int) (*entity.Category, error) {

	product, err := s.repositoryCategory.FindById(ID)

	if err != nil {
		return nil, err
	}

	if product.ID == 0 {
		return nil, err
	}

	return product, nil
}

func (s *serviceCategory) DeleteCategory(ID int) error {
	// Temukan produk berdasarkan ID
	product, err := s.repositoryCategory.FindById(ID)
	if err != nil {
		return err
	}

	// Hapus produk dari basis data
	_, err = s.repositoryCategory.Delete(product)
	if err != nil {
		return err
	}

	return nil
}

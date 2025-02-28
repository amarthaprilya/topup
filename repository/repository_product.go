package repository

import (
	"camera-rent/entity"

	"gorm.io/gorm"
)

type RepositoryProduct interface {
	FindAll() ([]*entity.Products, error)
	Save(product *entity.Products) (*entity.Products, error)
	FindById(ID int) (*entity.Products, error)
	Update(product *entity.Products) (*entity.Products, error)
	Delete(product *entity.Products) (*entity.Products, error)
	FindAllProductByCategory(ID int) ([]*entity.Products, error)
}

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db}
}

func (r *repositoryProduct) FindAllProductByCategory(ID int) ([]*entity.Products, error) {
	var product []*entity.Products

	err := r.db.Preload("Categorys").Where("category_id = ? ", ID).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryProduct) FindAll() ([]*entity.Products, error) {
	var product []*entity.Products

	err := r.db.Preload("Categorys").Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryProduct) Save(product *entity.Products) (*entity.Products, error) {
	err := r.db.Preload("Categorys").Create(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryProduct) FindById(ID int) (*entity.Products, error) {
	var product *entity.Products

	err := r.db.Preload("Categorys").Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryProduct) Update(product *entity.Products) (*entity.Products, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repositoryProduct) Delete(product *entity.Products) (*entity.Products, error) {
	err := r.db.Preload("Categorys").Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

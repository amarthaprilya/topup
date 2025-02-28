package repository

import (
	"camera-rent/entity"

	"gorm.io/gorm"
)

type RepositoryCategory interface {
	FindAll() ([]*entity.Category, error)
	Save(product *entity.Category) (*entity.Category, error)
	FindById(ID int) (*entity.Category, error)
	Update(product *entity.Category) (*entity.Category, error)
	Delete(product *entity.Category) (*entity.Category, error)
}

type repositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{db}
}

func (r *repositoryCategory) FindAll() ([]*entity.Category, error) {
	var product []*entity.Category

	err := r.db.Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryCategory) Save(product *entity.Category) (*entity.Category, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryCategory) FindById(ID int) (*entity.Category, error) {
	var product *entity.Category

	err := r.db.Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryCategory) Update(product *entity.Category) (*entity.Category, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repositoryCategory) Delete(product *entity.Category) (*entity.Category, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

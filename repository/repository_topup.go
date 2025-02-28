package repository

import (
	"camera-rent/entity"

	"gorm.io/gorm"
)

type RepositoryTopUp interface {
	Save(TopUp *entity.TopUp) (*entity.TopUp, error)
	FindById(ID int) (*entity.TopUp, error)
}

type repositoryTopUp struct {
	db *gorm.DB
}

func NewRepositoryTopUp(db *gorm.DB) *repositoryTopUp {
	return &repositoryTopUp{db}
}

func (r *repositoryTopUp) Save(topUp *entity.TopUp) (*entity.TopUp, error) {
	err := r.db.Preload("Users").Create(&topUp).Error

	if err != nil {
		return topUp, err
	}
	return topUp, nil
}

func (r *repositoryTopUp) FindById(ID int) (*entity.TopUp, error) {
	var topUp *entity.TopUp

	err := r.db.Preload("Users").Where("id = ?", ID).Find(&topUp).Error

	if err != nil {
		return topUp, err
	}
	return topUp, nil
}

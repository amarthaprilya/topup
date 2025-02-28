package repository

import (
	"camera-rent/entity"

	"gorm.io/gorm"
)

type RepositoryPaymentSaldo interface {
	FindAll() ([]*entity.PaymentSaldo, error)
	Save(payment *entity.PaymentSaldo) (*entity.PaymentSaldo, error)
	FindById(ID int) (*entity.PaymentSaldo, error)
	Update(payment *entity.PaymentSaldo) (*entity.PaymentSaldo, error)
	Delete(payment *entity.PaymentSaldo) error
	FindAllByUserID(ID int) ([]*entity.PaymentSaldo, error)
	FindByTransactionID(transactionID string) (*entity.PaymentSaldo, error)
	FindByOrderId(orderID string) (*entity.PaymentSaldo, error)
}

type repositoryPaymentSaldo struct {
	db *gorm.DB
}

func NewRepositoryPaymentSaldo(db *gorm.DB) *repositoryPaymentSaldo {
	return &repositoryPaymentSaldo{db}
}

func (r *repositoryPaymentSaldo) FindAll() ([]*entity.PaymentSaldo, error) {
	var payments []*entity.PaymentSaldo
	err := r.db.Preload("TopUps").Preload("Users").Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *repositoryPaymentSaldo) Save(payment *entity.PaymentSaldo) (*entity.PaymentSaldo, error) {
	if err := r.db.Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *repositoryPaymentSaldo) FindById(ID int) (*entity.PaymentSaldo, error) {
	var payment entity.PaymentSaldo
	err := r.db.Preload("TopUps").Preload("Users").First(&payment, ID).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *repositoryPaymentSaldo) Update(payment *entity.PaymentSaldo) (*entity.PaymentSaldo, error) {
	if err := r.db.Save(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *repositoryPaymentSaldo) Delete(payment *entity.PaymentSaldo) error {
	if err := r.db.Delete(payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositoryPaymentSaldo) FindAllByUserID(userID int) ([]*entity.PaymentSaldo, error) {
	var payments []*entity.PaymentSaldo
	err := r.db.Preload("TopUps").Preload("Users").Where("user_id = ?", userID).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *repositoryPaymentSaldo) FindByTransactionID(transactionID string) (*entity.PaymentSaldo, error) {
	var payment entity.PaymentSaldo
	err := r.db.Preload("TopUps").Preload("Users").Where("transaction_id = ?", transactionID).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *repositoryPaymentSaldo) FindByOrderId(orderID string) (*entity.PaymentSaldo, error) {
	var payment entity.PaymentSaldo
	err := r.db.Preload("TopUps").Preload("Users").Where("id = ?", orderID).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

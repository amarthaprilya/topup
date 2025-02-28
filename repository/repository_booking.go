package repository

import (
	"camera-rent/entity"
	"errors"

	"gorm.io/gorm"
)

type RepositoryBooking interface {
	FindAll() ([]*entity.Booking, error)
	Save(booking *entity.Booking) (*entity.Booking, error)
	FindById(ID int) (*entity.Booking, error)
	Update(booking *entity.Booking) (*entity.Booking, error)
	Delete(booking *entity.Booking) error
}

type repositoryBooking struct {
	db *gorm.DB
}

func NewRepositoryBooking(db *gorm.DB) *repositoryBooking {
	return &repositoryBooking{db}
}

func (r *repositoryBooking) FindAll() ([]*entity.Booking, error) {
	var bookings []*entity.Booking
	err := r.db.Preload("Products").Preload("Products.Categorys").Preload("Users").Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *repositoryBooking) Save(booking *entity.Booking) (*entity.Booking, error) {
	if err := calculateDaysRent(booking); err != nil {
		return nil, err
	}
	if err := r.db.Create(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *repositoryBooking) FindById(ID int) (*entity.Booking, error) {
	var booking entity.Booking
	err := r.db.Preload("Products").Preload("Products.Categorys").Preload("Users").First(&booking, ID).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *repositoryBooking) Update(booking *entity.Booking) (*entity.Booking, error) {
	if err := calculateDaysRent(booking); err != nil {
		return nil, err
	}
	if err := r.db.Save(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *repositoryBooking) Delete(booking *entity.Booking) error {
	if err := r.db.Delete(booking).Error; err != nil {
		return err
	}
	return nil
}

func calculateDaysRent(booking *entity.Booking) error {
	if booking.LastDateRent.Before(booking.FirstDateRent) {
		return errors.New("LastDateRent must be after FirstDateRent")
	}
	booking.DaysRent = int(booking.LastDateRent.Sub(booking.FirstDateRent).Hours() / 24)
	return nil
}

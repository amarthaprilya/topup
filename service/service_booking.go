package service

import (
	"camera-rent/entity"
	"camera-rent/input"
	"camera-rent/repository"
	"errors"
)

type ServiceBooking interface {
	GetAllBookings() ([]*entity.Booking, error)
	CreateBooking(userID int, input input.BookingInput) (*entity.Booking, error)
	GetBookingById(ID int) (*entity.Booking, error)
	// UpdateBooking(ID int, input entity.Booking) (*entity.Booking, error)
	DeleteBooking(ID int) error
}

type serviceBooking struct {
	repository        repository.RepositoryBooking
	repositoryProduct repository.RepositoryProduct
	repositoryUser    repository.RepositoryUser
}

func NewServiceBooking(repository repository.RepositoryBooking, repositoryProduct repository.RepositoryProduct, repositoryUser repository.RepositoryUser) *serviceBooking {
	return &serviceBooking{repository, repositoryProduct, repositoryUser}
}

func (s *serviceBooking) GetAllBookings() ([]*entity.Booking, error) {
	return s.repository.FindAll()
}

func (s *serviceBooking) CreateBooking(userID int, input input.BookingInput) (*entity.Booking, error) {
	if input.LastDateRent.Before(input.FirstDateRent) {
		return nil, errors.New("LastDateRent must be after FirstDateRent")
	}

	if input.ProductID == 0 {
		return nil, errors.New("productId tidak boleh kosong")
	}

	// Cek apakah kategori ada di database
	price, err := s.repositoryProduct.FindById(input.ProductID)
	if err != nil {
		return nil, errors.New("product tidak ditemukan")
	}

	findUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return nil, err
	}

	booking := &entity.Booking{}

	booking.ProductID = input.ProductID
	booking.UserID = findUser.ID
	booking.DaysRent = int(input.LastDateRent.Sub(input.FirstDateRent).Hours() / 24)
	booking.FirstDateRent = input.FirstDateRent
	booking.LastDateRent = input.LastDateRent
	booking.TotalPrice = price.RentCost * booking.DaysRent

	new, err := s.repository.Save(booking)
	if err != nil {
		return nil, err
	}

	return new, nil

}

// func (s *serviceBooking) UpdateBooking(ID int, input entity.Booking) (*entity.Booking, error) {
// 	booking, err := s.repository.FindById(ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Update fields
// 	booking.FirstDateRent = input.FirstDateRent
// 	booking.LastDateRent = input.LastDateRent
// 	booking.DaysRent = int(booking.LastDateRent.Sub(booking.FirstDateRent).Hours() / 24)
// 	booking.ProductID = input.ProductID
// 	booking.UserID = input.UserID

// 	return s.repository.Update(booking)
// }

func (s *serviceBooking) GetBookingById(ID int) (*entity.Booking, error) {
	return s.repository.FindById(ID)
}

func (s *serviceBooking) DeleteBooking(ID int) error {
	booking, err := s.repository.FindById(ID)
	if err != nil {
		return err
	}

	return s.repository.Delete(booking)
}

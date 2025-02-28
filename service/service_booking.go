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
	// Validasi tanggal sewa
	if input.LastDateRent.Before(input.FirstDateRent) {
		return nil, errors.New("LastDateRent must be after FirstDateRent")
	}

	// Validasi ProductID
	if input.ProductID == 0 {
		return nil, errors.New("productId tidak boleh kosong")
	}

	// Ambil data produk
	product, err := s.repositoryProduct.FindById(input.ProductID)
	if err != nil {
		return nil, errors.New("product tidak ditemukan")
	}

	// Validasi Quantity, misalnya harus lebih dari 0
	if input.Quantity <= 0 {
		return nil, errors.New("quantity harus lebih besar dari 0")
	}

	// Cek stok produk mencukupi
	if product.Stock < input.Quantity {
		return nil, errors.New("stok produk tidak mencukupi")
	}

	// Ambil data user
	findUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return nil, err
	}

	// Hitung jumlah hari sewa
	daysRent := int(input.LastDateRent.Sub(input.FirstDateRent).Hours() / 24)

	// Hitung total harga: RentCost x daysRent x Quantity
	totalPrice := product.RentCost * daysRent * input.Quantity

	// Validasi saldo user
	if findUser.Saldo < totalPrice {
		return nil, errors.New("saldo tidak mencukupi untuk menyewa produk ini")
	}

	// Buat entitas booking baru
	booking := &entity.Booking{
		ProductID:     input.ProductID,
		UserID:        findUser.ID,
		DaysRent:      daysRent,
		FirstDateRent: input.FirstDateRent,
		LastDateRent:  input.LastDateRent,
		Quantity:      input.Quantity,
		TotalPrice:    totalPrice,
	}

	// Simpan booking ke database
	newBooking, err := s.repository.Save(booking)
	if err != nil {
		return nil, err
	}

	// Kurangi saldo user
	findUser.Saldo -= totalPrice
	_, err = s.repositoryUser.Update(findUser)
	if err != nil {
		return nil, errors.New("gagal memperbarui saldo user setelah booking")
	}

	// Kurangi stok produk berdasarkan Quantity
	product.Stock -= input.Quantity
	_, err = s.repositoryProduct.Update(product)
	if err != nil {
		return nil, errors.New("gagal memperbarui stok produk setelah booking")
	}

	return newBooking, nil
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

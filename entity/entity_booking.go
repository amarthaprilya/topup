package entity

import "time"

type Booking struct {
	ID            int
	DaysRent      int
	FirstDateRent time.Time
	LastDateRent  time.Time
	TotalPrice    int
	ProductID     int
	Products      Products `gorm:"foreignKey:ProductID"`
	UserID        int
	Users         User `gorm:"foreignKey:UserID"`
	UpdatedAt     time.Time
	CreatedAt     time.Time
}

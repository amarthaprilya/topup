package entity

import "time"

type Products struct {
	ID          int
	Name        string
	RentCost    int
	Stock       int
	Description string
	CategoryID  int
	Categorys   Category `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

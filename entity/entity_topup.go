package entity

import "time"

type TopUp struct {
	ID        int
	Amount    int
	UserID    int
	Users     User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

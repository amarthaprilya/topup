package entity

import "time"

type PaymentSaldo struct {
	ID            int
	StatusPayment string
	TransactionID string
	TopUpID       int
	TopUps        TopUp `gorm:"foreignKey:TopUpID"`
	UserID        int
	Users         User `gorm:"foreignKey:UserID"`
	UpdatedAt     time.Time
	CreatedAt     time.Time
}
